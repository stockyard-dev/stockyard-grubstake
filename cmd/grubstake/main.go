package main
import ("fmt";"log";"net/http";"os";"github.com/stockyard-dev/stockyard-grubstake/internal/server";"github.com/stockyard-dev/stockyard-grubstake/internal/store")
func main(){port:=os.Getenv("PORT");if port==""{port="8800"};dataDir:=os.Getenv("DATA_DIR");if dataDir==""{dataDir="./grubstake-data"}
db,err:=store.Open(dataDir);if err!=nil{log.Fatalf("grubstake: %v",err)};defer db.Close();srv:=server.New(db)
fmt.Printf("\n  Grubstake — Self-hosted crowdfunding\n  Dashboard:  http://localhost:%s/ui\n  API:        http://localhost:%s/api\n\n",port,port)
log.Printf("grubstake: listening on :%s",port);log.Fatal(http.ListenAndServe(":"+port,srv))}
