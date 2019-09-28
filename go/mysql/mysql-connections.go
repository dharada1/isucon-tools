// https://gist.github.com/catatsuy/e627aaf118fbe001f2e7c665fda48146
// http://dsas.blog.klab.org/archives/2018-02/configure-sql-db.html
// https://www.alexedwards.net/blog/configuring-sqldb

db, _ = sql.Open("mysql", dsn)

maxConns := os.Getenv("DB_MAXOPENCONNS")
if maxConns != "" {
	i, err := strconv.Atoi(maxConns)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(i)
	db.SetMaxIdleConns(i)
}
