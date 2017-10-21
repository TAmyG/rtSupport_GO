0071CC81EBDB00000000000001

#Go Rethink db
go get gopkg.in/gorethink/gorethink.v3

##Crear BD
r.dbCreate('rtsupport')

r.db('rtsupport').table('channel').changes({
  includeInitial: true
})
  
##Crear tabla  
r.db('rtsupport').tableCreate('channel')
r.db('rtsupport').tableCreate('message')
r.db('rtsupport').tableCreate('user')


##Crear indices
r.db("rtsupport").table("channel").indexCreate("name")
r.db("rtsupport").table("user").indexCreate("name")
r.db("rtsupport").table("message").indexCreate("createdAt")