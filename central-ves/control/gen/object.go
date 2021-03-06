package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/HyperService-Consortium/go-ves/central-ves/model"
)

type ObjectCategories struct {
	artisan.VirtualService
	List    artisan.Category
	Post    artisan.Category
	Inspect artisan.Category
	IdGroup artisan.Category
}

func DescribeObjectService(base string) artisan.ProposingService {
	var objectModel = new(model.Object)
	svc := &ObjectCategories{
		List: artisan.Ink().
			Path("object-list").
			Method(artisan.POST, "ListObjects",
				artisan.QT("ListObjectsRequest", model.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("objects", objectModel)),
				),
			),
		Post: artisan.Ink().
			Path("object").
			Method(artisan.POST, "PostObject",
				artisan.Request(),
				artisan.Reply(
					codeField,
					artisan.Param("object", &objectModel),
				),
			),
		Inspect: artisan.Ink().Path("object/:oid/inspect").
			Method(artisan.GET, "InspectObject",
				artisan.Reply(
					codeField,
					artisan.Param("object", &objectModel),
				),
			),
		IdGroup: artisan.Ink().
			Path("object/:oid").
			Method(artisan.GET, "GetObject",
				artisan.Reply(
					codeField,
					artisan.Param("object", &objectModel),
				)).
			Method(artisan.PUT, "PutObject",
				artisan.Request()).
			Method(artisan.DELETE, "Delete"),
	}
	svc.Name("ObjectService").Base(base) //.
	//UseModel(artisan.Model(artisan.Name("object"), &objectModel))
	return svc
}
