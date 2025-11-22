/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_1103673869")

  // update collection data
  unmarshal({
    "deleteRule": "owner.id = @request.auth.id",
    "listRule": "owner.id = @request.auth.id",
    "updateRule": "owner.id = @request.auth.id",
    "viewRule": "owner.id = @request.auth.id"
  }, collection)

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_1103673869")

  // update collection data
  unmarshal({
    "deleteRule": "owner = @request.auth.id",
    "listRule": "owner = @request.auth.id",
    "updateRule": "owner = @request.auth.id",
    "viewRule": "owner = @request.auth.id"
  }, collection)

  return app.save(collection)
})
