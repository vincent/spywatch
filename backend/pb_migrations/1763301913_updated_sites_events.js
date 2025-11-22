/// <reference path="../pb_data/types.d.ts" />
migrate((app) => {
  const collection = app.findCollectionByNameOrId("pbc_2823905565")

  // update collection data
  unmarshal({
    "deleteRule": "resource.competitor.owner = @request.auth.id",
    "listRule": "resource.competitor.owner = @request.auth.id",
    "updateRule": "resource.competitor.owner = @request.auth.id",
    "viewRule": "resource.competitor.owner = @request.auth.id"
  }, collection)

  return app.save(collection)
}, (app) => {
  const collection = app.findCollectionByNameOrId("pbc_2823905565")

  // update collection data
  unmarshal({
    "deleteRule": "site.owner = @request.auth.id",
    "listRule": "site.owner = @request.auth.id",
    "updateRule": "site.owner = @request.auth.id",
    "viewRule": "site.owner = @request.auth.id"
  }, collection)

  return app.save(collection)
})
