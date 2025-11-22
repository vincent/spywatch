/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_1103673869")

  // add field
  collection.fields.addAt(5, new Field({
    "hidden": false,
    "id": "date952075636",
    "max": "",
    "min": "",
    "name": "last_diff_date",
    "presentable": false,
    "required": false,
    "system": false,
    "type": "date"
  }))

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_1103673869")

  // remove field
  collection.fields.removeById("date952075636")

  return app.save(collection)
})
