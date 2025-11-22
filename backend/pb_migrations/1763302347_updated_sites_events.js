/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_2823905565")

  // update collection data
  unmarshal({
    "name": "snapshots"
  }, collection)

  // remove field
  collection.fields.removeById("relation1766001124")

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_2823905565")

  // update collection data
  unmarshal({
    "name": "sites_events"
  }, collection)

  // add field
  collection.fields.addAt(1, new Field({
    "cascadeDelete": false,
    "collectionId": "pbc_1103673869",
    "hidden": false,
    "id": "relation1766001124",
    "maxSelect": 1,
    "minSelect": 0,
    "name": "site",
    "presentable": false,
    "required": false,
    "system": false,
    "type": "relation"
  }))

  return app.save(collection)
})
