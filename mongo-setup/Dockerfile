FROM mongo

COPY birds.json /birds.json

CMD mongoimport --uri mongodb://172.20.0.6:27017/birdai --db birdai --collection birds --file /birds.json --jsonArray
