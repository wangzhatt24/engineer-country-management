package engineer_implement

import (
	"context"
	"database/sql"
	pb "engineer-country-management/pkg/engineer/v1"
	"errors"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type EngineerImplement struct {
	pb.UnimplementedEngineerServiceServer
	DB *sql.DB
}

func (ei *EngineerImplement) AddEngineer(ctx context.Context, in *pb.AddEngineerRequest) (*pb.Engineer, error) {
	created_at, updated_at := time.Now(), time.Now()
	result, err := ei.DB.Exec("INSERT INTO engineer (first_name, last_name, gender, country_id, title, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		in.FirstName, in.LastName, in.Gender, in.CountryId, in.Title, created_at, updated_at,
	)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &pb.Engineer{
		Id:        id,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Gender:    in.Gender,
		CountryId: in.CountryId,
		Title:     in.Title,
		CreatedAt: timestamppb.New(created_at),
		UpdatedAt: timestamppb.New(updated_at),
	}, err
}

func (ei *EngineerImplement) GetEngineerById(ctx context.Context, in *pb.GetEngineerByIdRequest) (*pb.Engineer, error) {
	result := ei.DB.QueryRow("SELECT * FROM engineer WHERE id = ?", in.Id)

	var engineer pb.Engineer
	var created_at time.Time
	var updated_at time.Time

	err := result.Scan(&engineer.Id,
		&engineer.FirstName,
		&engineer.LastName,
		&engineer.Gender,
		&engineer.CountryId,
		&engineer.Title,
		&created_at,
		&updated_at)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("engineer not found")
		}

		return nil, err
	}

	// time convert
	engineer.CreatedAt = timestamppb.New(created_at)
	engineer.UpdatedAt = timestamppb.New(updated_at)

	return &engineer, nil
}

func (ei *EngineerImplement) DeleteEngineerById(ctx context.Context, in *pb.DeleteEngineerRequest) (*pb.Engineer, error) {
	engineer, err := ei.GetEngineerById(ctx, &pb.GetEngineerByIdRequest{Id: in.Id})
	if err != nil {
		return nil, errors.New("engineer not found")
	}

	result, err := ei.DB.Exec("DELETE FROM engineer WHERE id = ?", in.Id)
	if err != nil {
		return nil, err
	}

	ra, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if ra != 1 {
		return nil, errors.New("delete failed")
	}

	return engineer, nil
}

func (ei *EngineerImplement) UpdateEngineer(ctx context.Context, in *pb.UpdateEngineerRequest) (*pb.Engineer, error) {
	result, err := ei.DB.Exec("UPDATE engineer SET first_name = ?, last_name = ?, gender = ?, country_id = ?, title = ? WHERE id = ?",
		in.FirstName, in.LastName, in.Gender, in.CountryId, in.Title, in.Id,
	)

	if err != nil {
		return nil, err
	}

	ra, err := result.RowsAffected()
	if err != nil {
		return nil, errors.New("error when detect row effected")
	}
	if ra != 1 {
		return nil, errors.New("no row affected")
	}

	engineerAfterUpdated, err := ei.GetEngineerById(ctx, &pb.GetEngineerByIdRequest{Id: in.Id})
	if err != nil {
		return nil, err
	}

	return engineerAfterUpdated, nil
}

func (ei *EngineerImplement) ListEngineers(ctx context.Context, in *pb.ListEngineersRequest) (*pb.ListEngineersResponse, error) {
	pageNumber := in.GetPageNumber()
	pageSize := in.GetPageSize()

	// ensure that page number is valid
	if pageNumber <= 0 {
		pageNumber = 1 // default page number
	}

	// ensure page size is valid
	if pageSize <= 0 {
		pageSize = 10 // default page size
	}

	// cal offset, ex: 3 - 1 * 10 = 20, bat dau tu 20
	offset := (pageNumber - 1) * pageSize

	// total engineers
	var totalCount int64
	err := ei.DB.QueryRow("SELECT COUNT(*) FROM engineer").Scan(&totalCount)
	if err != nil {
		return nil, err
	}

	rows, err := ei.DB.Query("SELECT * FROM engineer LIMIT ? OFFSET ?", pageSize, offset)
	if err != nil {
		return nil, err
	}

	var engineers []*pb.Engineer

	for rows.Next() {
		var engineer pb.Engineer
		var created_at time.Time
		var updated_at time.Time

		err := rows.Scan(&engineer.Id,
			&engineer.FirstName,
			&engineer.LastName,
			&engineer.Gender,
			&engineer.CountryId,
			&engineer.Title,
			&created_at,
			&updated_at)

		if err != nil {
			return nil, err
		}

		// time convert
		engineer.CreatedAt = timestamppb.New(created_at)
		engineer.UpdatedAt = timestamppb.New(updated_at)

		engineers = append(engineers, &engineer)
	}

	return &pb.ListEngineersResponse{
		Engineers:  engineers,
		TotalCount: totalCount,
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}, nil
}
