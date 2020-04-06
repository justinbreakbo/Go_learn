package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TLessonController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TLessonController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TLessonController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TLessonController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TLessonController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TLessonController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TLessonController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TLessonController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TLessonController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TLessonController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TMatchController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TMatchController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TMatchController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TMatchController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TMatchController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TMatchController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TMatchController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TMatchController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TMatchController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TMatchController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOptionsController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOptionsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOptionsController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOptionsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOptionsController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOptionsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOptionsController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOptionsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOptionsController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOptionsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOrganizationController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOrganizationController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOrganizationController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOrganizationController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOrganizationController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOrganizationController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOrganizationController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOrganizationController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOrganizationController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TOrganizationController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TProjectController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TProjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TProjectController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TProjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TProjectController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TProjectController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TProjectController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TProjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TProjectController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TProjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TStuLessonController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TStuLessonController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TStuLessonController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TStuLessonController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TStuLessonController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TStuLessonController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TStuLessonController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TStuLessonController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TStuLessonController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TStuLessonController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TUserMatchController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TUserMatchController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TUserMatchController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TUserMatchController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TUserMatchController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TUserMatchController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TUserMatchController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TUserMatchController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TUserMatchController"] = append(beego.GlobalControllerRouter["Go_learn/beego_learn_separate_ogc/controllers:TUserMatchController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
