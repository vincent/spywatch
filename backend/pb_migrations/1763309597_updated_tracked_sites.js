/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_1103673869")

  // update collection data
  unmarshal({
    "name": "competitors"
  }, collection)

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_1103673869")

  // update collection data
  unmarshal({
    "name": "tracked_sites"
  }, collection)

  return app.save(collection)
})
