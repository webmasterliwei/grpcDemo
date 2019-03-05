package main

import (
	"grpcDemo/mr_gun"
	"golang.org/x/net/context"
	_ "github.com/go-sql-driver/mysql"
	"net"
	"log"
	"google.golang.org/grpc"
	"database/sql"
	"errors"
)

const (
	port = ":50051"
)

var db *sql.DB
var sqlMap = map[string]string{
	"1": "select * from gun order by total_damage desc",
	"2": "select * from gun where scatter = 0 order by total_damage desc",
	"3": "select * from gun where scatter = 1 order by total_damage desc",
	"4": "select * from gun where bullets_shot = 1 order by total_damage desc",
	"5": "select * from gun order by bullets_shot desc,total_damage desc",
}

func init() {
	var err error
	db, err = sql.Open("mysql", "root:liwei417423@tcp(127.0.0.1:3306)/mr_gun?charset=utf8")
	if err != nil {
		log.Fatalf("mysql connect failed: %v", err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	if err = db.Ping(); err != nil {
		log.Fatalf("mysql connect failed: %v", err)
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	mr_gun.RegisterMrGunServer(s, &MrGunServer{})
	s.Serve(lis)
}

type MrGunServer struct{}

func (s *MrGunServer) GetRank(ctx context.Context, request *mr_gun.Request) (*mr_gun.Reply, error) {
	var gunInfo []*mr_gun.GunInfo
	Sql, found := sqlMap[request.Sort]
	if !found {
		return nil, errors.New("sort error")
	}
	rows, err := db.Query(Sql)
	if err != nil {
		return nil, err
	} else {
		defer rows.Close()
	}
	for rows.Next() {
		var gun mr_gun.GunInfo
		err = rows.Scan(&gun.Name, &gun.Damage, &gun.BulletsShot, &gun.TotalDamage, &gun.Scatter)
		if err != nil {
			return nil, err
		}
		gunInfo = append(gunInfo, &gun)
	}
	reply := mr_gun.Reply{GunInfo: gunInfo}
	return &reply, nil
}
