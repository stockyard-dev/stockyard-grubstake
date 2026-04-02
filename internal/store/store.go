package store
import ("database/sql";"fmt";"os";"path/filepath";"time";_ "modernc.org/sqlite")
type DB struct{db *sql.DB}
type Campaign struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	GoalCents int `json:"goal_cents"`
	RaisedCents int `json:"raised_cents"`
	BackerCount int `json:"backer_count"`
	Status string `json:"status"`
	EndsAt string `json:"ends_at"`
	CreatedAt string `json:"created_at"`
}
func Open(d string)(*DB,error){if err:=os.MkdirAll(d,0755);err!=nil{return nil,err};db,err:=sql.Open("sqlite",filepath.Join(d,"grubstake.db")+"?_journal_mode=WAL&_busy_timeout=5000");if err!=nil{return nil,err}
db.Exec(`CREATE TABLE IF NOT EXISTS campaigns(id TEXT PRIMARY KEY,name TEXT NOT NULL,description TEXT DEFAULT '',goal_cents INTEGER DEFAULT 0,raised_cents INTEGER DEFAULT 0,backer_count INTEGER DEFAULT 0,status TEXT DEFAULT 'active',ends_at TEXT DEFAULT '',created_at TEXT DEFAULT(datetime('now')))`)
return &DB{db:db},nil}
func(d *DB)Close()error{return d.db.Close()}
func genID()string{return fmt.Sprintf("%d",time.Now().UnixNano())}
func now()string{return time.Now().UTC().Format(time.RFC3339)}
func(d *DB)Create(e *Campaign)error{e.ID=genID();e.CreatedAt=now();_,err:=d.db.Exec(`INSERT INTO campaigns(id,name,description,goal_cents,raised_cents,backer_count,status,ends_at,created_at)VALUES(?,?,?,?,?,?,?,?,?)`,e.ID,e.Name,e.Description,e.GoalCents,e.RaisedCents,e.BackerCount,e.Status,e.EndsAt,e.CreatedAt);return err}
func(d *DB)Get(id string)*Campaign{var e Campaign;if d.db.QueryRow(`SELECT id,name,description,goal_cents,raised_cents,backer_count,status,ends_at,created_at FROM campaigns WHERE id=?`,id).Scan(&e.ID,&e.Name,&e.Description,&e.GoalCents,&e.RaisedCents,&e.BackerCount,&e.Status,&e.EndsAt,&e.CreatedAt)!=nil{return nil};return &e}
func(d *DB)List()[]Campaign{rows,_:=d.db.Query(`SELECT id,name,description,goal_cents,raised_cents,backer_count,status,ends_at,created_at FROM campaigns ORDER BY created_at DESC`);if rows==nil{return nil};defer rows.Close();var o []Campaign;for rows.Next(){var e Campaign;rows.Scan(&e.ID,&e.Name,&e.Description,&e.GoalCents,&e.RaisedCents,&e.BackerCount,&e.Status,&e.EndsAt,&e.CreatedAt);o=append(o,e)};return o}
func(d *DB)Delete(id string)error{_,err:=d.db.Exec(`DELETE FROM campaigns WHERE id=?`,id);return err}
func(d *DB)Count()int{var n int;d.db.QueryRow(`SELECT COUNT(*) FROM campaigns`).Scan(&n);return n}
