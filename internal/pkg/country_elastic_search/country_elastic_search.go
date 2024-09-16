package country_elastic_search

import (
	"context"

	"engineer-country-management/internal/pkg/queues/elastic_search/elastic_search_queue_constants"
	pb "engineer-country-management/pkg/country/v1"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"google.golang.org/protobuf/encoding/protojson"
)

func SearchCountryFuzzyByName(ctx context.Context, esClient *elasticsearch.TypedClient, countryName string) (*pb.Countries, error) {
	res, err :=
		esClient.
			Search().
			Index(elastic_search_queue_constants.COUNTRY_INDEX).
			Request(&search.Request{
				Query: &types.Query{
					Bool: &types.BoolQuery{
						Should: []types.Query{
							{
								MatchPhrasePrefix: map[string]types.MatchPhrasePrefixQuery{
									"country_name": {
										Query: countryName,
									},
								},
							},
							{
								Fuzzy: map[string]types.FuzzyQuery{
									"country_name": {
										Value:     countryName,
										Fuzziness: "auto",
									},
								},
							},
						},
					},
				},
			}).
			Do(context.TODO())

	if err != nil {
		return nil, err
	}

	var countries pb.Countries

	for _, hit := range res.Hits.Hits {
		var country pb.Country
		err := protojson.Unmarshal(hit.Source_, &country)
		if err != nil {
			return nil, err
		}
		countries.Country = append(countries.Country, &country)
	}

	return &countries, nil
}
