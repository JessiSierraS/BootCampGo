package products

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

//Controller / Estructura
type MySQLRepository struct{
	Repository
	//base de datos al puntero DB
	Database *sql.DB
}

//Constructor
func NewRepositorySQL(db *sql.DB) Repository {
	return &MySQLRepository{Database: db}
}

func (repository *MySQLRepository) Get(id int) (product *Product, err error){
	//Query product.
	query := `
		SELECT
			id, name, quantity, code_value, is_published, expiration, price
		FROM products
		WHERE
			id = ?;
		`
	row := repository.Database.QueryRow(query,id)
	if err != nil {
		return
	}

	//Check if product exists.
	if row.Err() != nil {
		switch row.Err() {
		case sql.ErrNoRows:
			err = ErrNotFound
		default:
			err = ErrInternal
		}
		return
	}

	//Scan product
	s := Product{}
	err = row.Scan(
		&s.ID,
		&s.Name,
		&s.Quantity,
		&s.Code_value,
		&s.Is_published,
		&s.Expiration,
		&s.Price,
	)

	product = &s
	if err != nil{
		err = ErrInternal
		return
 	}

	//Everything is ok.
	return 
}

func (repository *MySQLRepository) Store(product *Product) (err error) {
	//Insert product
	statement, err := repository.Database.Prepare(`
		INSERT INTO products (
			name, quantity, code_value, is_published, expiration, price
		) VALUES (
			?, ?, ?, ?, ?, ?
		);
	`)
	if err != nil {
		err = ErrInternal
		return
	}
	defer statement.Close()

	result, err := statement.Exec(
		product.Name,
		product.Quantity,
		product.Code_value,
		product.Is_published,
		product.Expiration,
		product.Price,
	)
	if err != nil {
		driverErr, ok := err.(*mysql.MySQLError)
		if !ok {
			err = ErrInternal
			return
		}

		switch driverErr.Number{
		case 1862:
			err = ErrDuplicated
		default:
			err = ErrInternal
		}
		return
	}

	//Check if product was inserted
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected  != 1{
		err = ErrInternal
		return
	}

	//Get product id.
	productID, err := result.LastInsertId()
	if err != nil {
		err = ErrInternal
		return
	}

	//Everrything is ok.
	product.ID = int(productID)
	return
}

func (repository *MySQLRepository) Update(product *Product) (err error){
	statement, err := repository.Database.Prepare(
		`UPDATE products 
		SET name = ?, quantity= ?, code_value= ?, is_published= ?, expiration= ?, price= ?
		WHERE id =?; `)
	if err != nil{
		err = ErrInternal
		return 
	}
	defer statement.Close()

	result, err := statement.Exec(
		product.Name,
		product.Quantity,
		product.Code_value,
		product.Is_published,
		product.Expiration,
		product.Price,
		product.ID,
	)

	_ , err = result.RowsAffected()

	if err != nil{
		driverErr, ok := err.(*mysql.MySQLError)
		if !ok {
			err = ErrInternal
			return 
		}

		switch driverErr.Number{
		case 1862:
			err = ErrDuplicated
		default:
			err = ErrInternal
		}
	}
	return
}

func (repository *MySQLRepository) Delete(id int) (err error) {
	//Delete product
	statement, err := repository.Database.Prepare(`
		DELETE FROM products
		WHERE
			id = ?;
	`)
	if err != nil {
		err = ErrInternal
		return
	}
	
	result, err := statement.Exec(id)
	if err != nil {
		//TODO: validar según error code
		err = ErrInternal
		return
	}

	//Check if product was deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		err = ErrInternal
		return
	}

	if rowsAffected == 0 {
		err = ErrNotFound
		return
	}
	//Everrything is ok.
	return
}