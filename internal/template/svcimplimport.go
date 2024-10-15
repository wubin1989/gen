package template

const SvcImplImportGrpc = `
	"context"
	"github.com/unionj-cloud/toolkit/errorx"
	"{{.ConfigPackage}}"
	"{{.DtoPackage}}"
	"{{.ModelPackage}}"
	"{{.QueryPackage}}"
	"github.com/unionj-cloud/toolkit/copier"
	"github.com/unionj-cloud/go-doudou/v2/framework/database"
	paginate "github.com/unionj-cloud/toolkit/pagination/gorm"
	"github.com/pkg/errors"
	pb "{{.TransportGrpcPackage}}"
`

const SvcImplImportRest = `
	"context"
	"github.com/unionj-cloud/toolkit/errorx"
	"{{.ConfigPackage}}"
	"{{.DtoPackage}}"
	"{{.ModelPackage}}"
	"{{.QueryPackage}}"
	"github.com/unionj-cloud/toolkit/copier"
	"github.com/unionj-cloud/go-doudou/v2/framework/database"
	paginate "github.com/unionj-cloud/toolkit/pagination/gorm"
	"github.com/pkg/errors"
`
