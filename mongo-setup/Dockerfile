FROM mongo

COPY birds.json /birds.json
COPY media.json /media.json

CMD mongoimport --uri mongodb://mongodb:27017/birdai --db birdai --collection birds --file /birds.json --jsonArray && \
mongoimport --uri mongodb://mongodb:27017/birdai --db birdai --collection media --file /media.json --jsonArray
