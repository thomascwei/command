import pymongo
import json

myclient = pymongo.MongoClient(host="mongodb://thomas:123456@localhost/", port=27017)

mydb = myclient["command"]
mycol = mydb["command_template"]

with open('test_data/temp.json', encoding="UTF-8") as f:
    data = json.load(f)
    mycol.insert_many(data)

mycol = mydb["header_template"]
with open('test_data/header.json', encoding="UTF-8") as f:
    data = json.load(f)
    mycol.insert_many(data)
