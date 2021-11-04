# Process
## Install PostgreSQL
## Create a User and Pwd with superuser rights
## Create a table with following columns:
```
isbn    character(14)          serial  primary key  not null
title   text                                        not null
author  character varying(255)                      not null
price   numeric(5,2)             default value: 20.00       not null
```

## Update the password and user fields

Remark : don't forget to convert the price into float before recording in the DB if you use the templates
```
f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		http.Error(w, http.StatusText(406)+"Please hit back and enter a number for the price", http.StatusNotAcceptable)
		return
	}
	bk.Price = float32(f64)
```
