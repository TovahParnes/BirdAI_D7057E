import http.client
import json
from bson import ObjectId
import urllib.parse
import os
import re


def get_wikipedia_url(bird_name):
    normalized_bird_name = urllib.parse.quote(bird_name.title().encode("utf-8"))
    url = "en.wikipedia.org"
    endpoint = f"/w/api.php?action=query&list=search&format=json&srsearch={normalized_bird_name}"

    conn = http.client.HTTPSConnection(url)

    try:
        conn.request("GET", endpoint)
        response = conn.getresponse()

        # Check if OK
        if response.status == 200:
            response_data = response.read()
            data = json.loads(response_data.decode("utf-8"))

            # Check if there are search results
            if "query" in data and "search" in data["query"] and data["query"]["search"]:
                page_title = urllib.parse.quote(data["query"]["search"][0]["title"].replace(" ", "_"))
                return f"https://en.wikipedia.org/wiki/{page_title}"
            else:
                print(f"No search results for {bird_name}")
                return None
        else:
            print(f"Request failed with status code: {response.status}")
            return None

    finally:
        conn.close()


def get_audio(bird_name):
    normalized_bird_name = urllib.parse.quote(bird_name.title().encode("utf-8"))
    normalized_bird_name = normalized_bird_name.replace("%20", "+")
    url = "xeno-canto.org"
    endpoint = f"/api/2/recordings?query={normalized_bird_name}"

    conn = http.client.HTTPSConnection(url)

    try:
        conn.request("GET", endpoint)
        response = conn.getresponse()

        # Check if OK
        if response.status == 200:
            response_data = response.read()
            data = json.loads(response_data.decode("utf-8"))
            if int(data.get("numRecordings")) > 0:
                return data.get("recordings")[0].get("file")
            return

        else:
            print(f"Request failed with status code: {response.status}")
            return

    finally:
        conn.close()


def get_bird_name(wikipedia_url):
    match = re.search(r"/wiki/(.+)$", wikipedia_url)
    if match:
        return match.group(1).replace("_", " ")
    return None


def print_progress(i, total, size):
    percent = 100 * (i / float(total))
    completed = int(size * i // total)
    bar = "#" * completed + "-" * (size - completed)
    print(f"Progress: |{bar}| {percent:.2f}% ({i}/{total})")


def load_json_file(file_path):
    json_data = []
    try:
        with open(file_path, "r", encoding="utf-8") as file:
            json_data = json.load(file)
    except Exception as e:
        print(f"Error loading from json: {e}")
    return json_data


def save_json_file(file_path, json_data):
    try:
        with open(file_path, "w", encoding="utf-8") as file:
            json.dump(json_data, file, indent=2)
    except Exception as e:
        print(f"Error writing to birds.json: {e}")


def process_text_file(directory, filename, all_birds, json_data, species):
    file_path = os.path.join(directory, filename)
    try:
        with open(file_path, "r") as file:
            for line in file:
                name = line.strip()
                if name not in all_birds:
                    new_id = str(ObjectId())
                    bird_dict = {"_id": new_id, "name": name, "species": species}
                    all_birds[name] = bird_dict
                    json_data.append(bird_dict)
    except Exception as e:
        print(f"Error loading from text file '{filename}': {e}")


def main():
    current_directory = os.getcwd()
    json_data = load_json_file(os.path.join(current_directory, "birds.json"))
    all_birds = {entry["name"]: entry for entry in json_data}

    process_text_file(current_directory, "birds.txt", all_birds, json_data, False)
    process_text_file(current_directory, "genus.txt", all_birds, json_data, True)

    errors = []

    for i, entry in enumerate(json_data):
        if i % 10 == 0 or i == len(json_data) - 1:
            print_progress(i, len(json_data), 40)
        bird_name = entry.get("name")
        if bird_name:
            wikipedia_url = entry.get("description") or get_wikipedia_url(bird_name)
            audio_url = entry.get("sound_id") or get_audio(bird_name)

            # Wikipedia url checks
            if entry.get("description") is None and wikipedia_url:
                entry["description"] = wikipedia_url
                print(f"{wikipedia_url:55} was added to {bird_name}")
            elif entry.get("description") and wikipedia_url:
                continue
            elif entry.get("description") is None and wikipedia_url is None:
                errors.append(f"Couldn't find wikipeida url for {bird_name}")

            # Xeno-canto audio url checks
            if entry.get("sound_id") is None and audio_url:
                entry["sound_id"] = audio_url
                print(f"{audio_url:55} was added to {bird_name}")
            elif entry.get("sound_id") and audio_url:
                continue
            elif entry.get("sound_id") is None and audio_url is None:
                if wikipedia_url:
                    get_audio(get_wikipedia_url(bird_name))
                    if audio_url:
                        entry["sound_id"] = audio_url
                        print(f"{audio_url:55} was added to {bird_name}")
                    else:
                        errors.append(f"Couldn't find audio url for {bird_name}")

    print_progress(len(json_data), len(json_data), 40)
    # Save the json_data after processing
    file_path = os.path.join(current_directory, "birds.json")
    save_json_file(file_path, json_data)

    if errors:
        for err in errors:
            print(err)
        print("Add error ones manually if correct results can be found.")


if __name__ == "__main__":
    main()
