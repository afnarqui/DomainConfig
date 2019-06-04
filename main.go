package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/go-chi/chi"
	"time"
	"strings"
	"os"
	"io/ioutil"
	"encoding/json"
	"errors"
)

var Host string
var db *sql.DB
var domainnew = Domain{}
var domainold = Domain{}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	
	log.Println("Corriendo en http://localhost:8001")
	r := chi.NewRouter()

	r.Get("/public", func(w http.ResponseWriter, r *http.Request) {

		nombre := r.URL.Query().Get("nombre")
		
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		url := "https://api.ssllabs.com/api/v3/analyze?host="+nombre
		        
		response, err := http.Get(url)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		j := "["+string(responseData)+"]"
		xp := []Domain{}
	
		errr := json.Unmarshal([]byte(j), &xp)
		if errr != nil {
			fmt.Println(errr)
		}
		data := Domain{}
		endpointsssf := Endpointss{}

		for i, v := range xp {
			 fmt.Println(i,v)
			 data.Host = v.Host
			 Host = v.Host
			 data.Port = v.Port
			 data.Protocol = v.Protocol
			 data.IsPublic = v.IsPublic
			 data.Status = v.Status
			
			for b, k := range v.Endpoints {
				endpointsss := Endpointss{
					Endpoints{
						Grade:k.Grade,
						IpAddress:k.IpAddress,
						ServerName : k.ServerName,
						StatusMessage : k.StatusMessage,
						GradeTrustIgnored : k.GradeTrustIgnored,
						HasWarnings : k.HasWarnings,
						IsExceptional : k.IsExceptional,
						Progress : k.Progress,
						Duration : k.Duration,
						Delegation : k.Delegation,
					},
				}
	
				fmt.Println(b)
				endpointsssf = endpointsss
			}
			data.Endpoints = endpointsssf 
		}
		domainnew = data
	
		n := new(Domain)
		domain, err := n.GetAllDomain()
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		
		dataDomain, err := json.Marshal(domain)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		responsedata:= []Domain{}
	
		errrs := json.Unmarshal([]byte(dataDomain), &responsedata)
		if errrs != nil {
			fmt.Println(errrs)
		}
		fmt.Println(responsedata)
		if len(responsedata) > 0 {
			var dataupdate Domain
			domainold = responsedata[0]
			fmt.Println("should update")
			err = dataupdate.DeleteDomain()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}		
		} else {
			var datanew Domain
			fmt.Println("should add")
			err = datanew.CreateDomain()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		json.NewEncoder(w).Encode(domainnew)
})
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	direccion := ":8081" 
	fmt.Println("Server listing in " + direccion)

	log.Fatal(http.ListenAndServe(direccion+"/public/index.html", nil))
	
}

func Logger() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println(time.Now(), r.Method, r.URL)
        router.ServeHTTP(w, r) 
    })
}
var router *chi.Mux

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

func GetConnection() *sql.DB {
	if db != nil {
		return db
	}

	var err error

	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")

	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	return db
}

type Domain struct {
	Host            string      `json:"host"`
	Port            int         `json:"port"`
	Protocol        string      `json:"protocol"`
	IsPublic        bool        `json:"isPublic"`
	Status          string      `json:"status"`
	Endpoints       []Endpoints `json:"endpoints"`
}

type Endpoints struct {
	IpAddress         string `json:"ipAddress"`
	ServerName        string `json:"serverName"`
	StatusMessage     string `json:"statusMessage"`
	Grade             string `json:"grade"`
	GradeTrustIgnored string `json:"gradeTrustIgnored"`
	HasWarnings       bool   `json:"hasWarnings"`
	IsExceptional     bool   `json:"isExceptional"`
	Progress          int    `json:"progress"`
	Duration          int    `json:"duration"`
	Delegation        int    `json:"delegation"`
}

type Endpointss []Endpoints

func (n *Domain) GetAllDomain() ([]Domain, error) {
	db := GetConnection()
	Host = "'"+Host+"'"

	q := "select distinct host,port,protocol,ispublic,status from domain where host="+string(Host)
	rows, err := db.Query(q)
	if err != nil {
		return []Domain{}, err
	}
	defer rows.Close()
	bks := make([]Domain, 0)
	for rows.Next() {
		bk := Domain{}
		err := rows.Scan(&bk.Host, &bk.Port,&bk.Protocol,&bk.IsPublic,&bk.Status) 
		if err != nil {
			panic(err)
		}
		bks = append(bks, bk)
	}
	return bks, nil 
}

func (n Domain) CreateDomain() error {

	var host = domainnew.Host
	var port = domainnew.Port
	var protocol = domainnew.Protocol 
	var isPublic = domainnew.IsPublic
	fmt.Println("IsPublic")
	fmt.Println(domainnew.IsPublic)
	var status = domainnew.Status
	
	q := "INSERT INTO domain(host,port,protocol,ispublic,status) VALUES ($1,$2,$3,$4,$5)"
		db := GetConnection()
		defer db.Close()
		fmt.Println("should save")
		fmt.Println(q)
		stmt, err := db.Prepare(q)

		if err != nil {
		return err
		}
		defer stmt.Close()
	
		r, err := stmt.Exec(host,port,protocol,isPublic,status)
		if err != nil {
		return err
		}

		i, _ := r.RowsAffected()

		if i != 1 {
		return errors.New("Should error rows")
		}


	return nil
}

func (n Domain) DeleteDomain() error {

	dbdomain := GetConnection()

	var host = domainnew.Host

	qdomain := `DELETE FROM domain
		WHERE host=$1`
	stmtdomain, errdomain := dbdomain.Prepare(qdomain)
	if errdomain != nil {
		return errdomain
	}
	defer stmtdomain.Close()

	rdomain, errdomain := stmtdomain.Exec(host)
	if errdomain != nil {
		return errdomain
	}
	if idomain, errdomain := rdomain.RowsAffected(); errdomain != nil || idomain != 1 {
		return errors.New("fatal errors")
	}

	var portnewdomain = domainnew.Port
	var protocolnewdomain = domainnew.Protocol 
	var isPublicnewdomain = domainnew.IsPublic 
	var statusnewdomain = domainnew.Status
	qnewdomain := `INSERT INTO 
	domain(host,port,protocol,ispublic,status)
	VALUES ($1,$2,$3,$4,$5)`
     
		dbnewdomain := GetConnection()
		defer dbnewdomain.Close()

		stmtnewdomain, errnewdomain := dbnewdomain.Prepare(qnewdomain)

		if errnewdomain != nil {
		return errnewdomain
		}
		defer stmtnewdomain.Close()
		
		rnewdomain, errnewdomain := stmtnewdomain.Exec(host,portnewdomain,protocolnewdomain,isPublicnewdomain,statusnewdomain)
		if errnewdomain != nil {
		return errnewdomain
		}

		inewdomain, _ := rnewdomain.RowsAffected()

		if inewdomain != 1 {
		return errors.New("Should error rows newdomain")
		}

	return nil
}

