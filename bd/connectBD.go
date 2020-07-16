package bd

import (
    "context"
    "log"


    "go.mongo.org/mongo-driver/mongo"
    "go.mongo.org/mongo-driver/mongo/options"
)
//MongoC es el objeto de conexion a la base de datos
var MongoC=conectarBD()
var clienteOptions=options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster0.pjwyc.mongodb.net/<dbname>?retryWrites=true&w=majority")

/*ConnectarBd es la funcion que permite conectar con la base de datos*/
func ConnectarBd()*mongo.Client{
    client, err:=mongo.Connect(context.TODO(),clienteOptions)
    if err!=nil{
        log.Fatal(err.Error())
        return client
    }
    err = client.Ping(context.TODO(), nil)
    if err!= nil{
        log.Fatal(err.Error())
        return client
    }
    log.Println("Conexion exitosa con la Bd")
    return client
}
/*ChequeoConexion es el ping a la base de datos*/

func ChequeoConexion()int{
    err:= MongoC.Ping(context.TODO(), nil)
    if err!= nil{
        return 0
    }
    return 1
}