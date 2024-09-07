package country

import (
	"context"
	"engineer-country-management/internal/pkg/cache"
	"fmt"
	"time"

	mysqlWrapper "engineer-country-management/internal/pkg/mysql"
	pb "engineer-country-management/pkg/country/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var db = mysqlWrapper.GetClient()

// mysql sections
func MysqlFetchCountryById(id int64) (*pb.Country, error) {
	row := db.QueryRow("SELECT * FROM country WHERE id = ?", id)

	var country pb.Country
	var created_at time.Time
	var updated_at time.Time

	err := row.Scan(&country.Id, &country.CountryName, &created_at, &updated_at)
	if err != nil {
		return nil, err
	}

	return &pb.Country{
		Id:          country.Id,
		CountryName: country.CountryName,
		CreatedAt:   timestamppb.New(created_at),
		UpdatedAt:   timestamppb.New(updated_at),
	}, nil
}

func MysqlAddCountry(ctx context.Context, in *pb.AddCountryRequest) (*pb.Country, error) {
	created_at, updated_at := time.Now(), time.Now()
	result, err := db.Exec("INSERT INTO country(country_name, created_at, updated_at) VALUES (?, ?, ?)", in.GetCountryName(), created_at, updated_at)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &pb.Country{
		Id:          id,
		CountryName: in.CountryName,
		CreatedAt:   timestamppb.New(created_at),
		UpdatedAt:   timestamppb.New(updated_at),
	}, nil
}

func MysqlDeleteCountry(ctx context.Context, in *pb.DeleteCountryRequest) (*pb.Country, error) {
	// country, err := s.GetCountryById(ctx, &pb.GetCountryRequest{Id: in.Id})
	country, err := MysqlFetchCountryById(in.GetId())
	if err != nil {
		return nil, fmt.Errorf("\nerror when deleting country (mysqlDeleteCountry) %v", err)
	}

	result, err := db.Exec("DELETE FROM country where id = ?", in.GetId())

	if err != nil {
		return nil, fmt.Errorf("\nerror when delete country %v", err)
	}

	ra, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("\nerror checking row affected %v", err)
	}

	if ra == 1 {
		return country, nil
	}

	return nil, fmt.Errorf("\nerror when deleting country")
}

func MysqlUpdateCountry(ctx context.Context, in *pb.UpdateCountryRequest) (*pb.Country, error) {
	result, err := db.Exec("UPDATE country SET country_name = ? WHERE country.id = ?", in.GetCountryName(), in.GetId())

	if err != nil {
		return nil, err
	}

	rf, err := result.RowsAffected()
	if err != nil {
		fmt.Println("\nerror when check row affected")
	}

	if rf != 1 {
		fmt.Println("\nno rows updated")
	}

	country, err := MysqlFetchCountryById(in.GetId())
	if err != nil {
		return nil, fmt.Errorf("error when fetch country by id (mysqlFetchCountryById): %v", err)
	}

	err = cache.RedisUpdateCountryById(ctx, country)
	if err != nil {
		fmt.Printf("\nerror when update country to redis %v", err)
	}

	return country, nil
}

func MysqlListCountries(in *pb.ListCountriesRequest) (*pb.ListCountriesResponse, error) {
	// get total count
	// get records
	pageNumber := in.GetPageNumber()
	pageSize := in.GetPageSize()

	// ensure that page number is valid
	if pageNumber <= 0 {
		pageNumber = 1 // default
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (pageNumber - 1) * pageSize

	// total countries
	var totalCount int64
	err := db.QueryRow("SELECT COUNT(*) FROM country").Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM country LIMIT ? OFFSET ?", pageSize, offset)
	if err != nil {
		return nil, err
	}

	var countries pb.Countries

	for rows.Next() {
		var country pb.Country
		var created_at time.Time
		// var updated_at time.Time

		err := rows.Scan(&country.Id, &country.CountryName, &created_at, &created_at)
		if err != nil {
			return nil, err
		}

		countries.Country = append(countries.Country, &country)
	}

	return &pb.ListCountriesResponse{
		Countries:  &countries,
		TotalCount: totalCount,
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}, nil
}

// end mysql sections
