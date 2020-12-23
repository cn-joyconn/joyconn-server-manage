package controllers

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	config "openvpn-server-manage/go-openvpn/client/config"

	
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	lib "openvpn-server-manage/lib"
	models "openvpn-server-manage/models"
)

type NewCertParams struct {
	Name string `form:"Name" valid:"Required;"`
}

type CertificatesController struct {
	BaseController
}

func (c *CertificatesController) NestPrepare() {
	if !c.IsLogin {
		c.Ctx.Redirect(302, c.LoginPath())
		return
	}
	c.Data["breadcrumbs"] = &BreadCrumbs{
		Title: "Certificates",
	}
}

// @router /certificates/:key [get]
func (c *CertificatesController) Download() {
	name := c.GetString(":key")
	filename := fmt.Sprintf("%s.zip", name)

	c.Ctx.Output.Header("Content-Type", "application/zip")
	c.Ctx.Output.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

	zw := zip.NewWriter(c.Controller.Ctx.ResponseWriter)

	keysPath := models.GlobalCfg.OVConfigPath + "keys/"
	if cfgPath, err := saveClientConfig(name); err == nil {
		addFileToZip(zw, cfgPath)
	}
	addFileToZip(zw, keysPath+"ca.crt")
	addFileToZip(zw, keysPath+name+".crt")
	addFileToZip(zw, keysPath+name+".key")

	if err := zw.Close(); err != nil {
		logs.Error(err)
	}
}

func addFileToZip(zw *zip.Writer, path string) error {
	header := &zip.FileHeader{
		Name:         filepath.Base(path),
		Method:       zip.Store,
		ModifiedTime: uint16(time.Now().UnixNano()),
		ModifiedDate: uint16(time.Now().UnixNano()),
	}
	fi, err := os.Open(path)
	if err != nil {
		logs.Error(err)
		return err
	}

	fw, err := zw.CreateHeader(header)
	if err != nil {
		logs.Error(err)
		return err
	}

	if _, err = io.Copy(fw, fi); err != nil {
		logs.Error(err)
		return err
	}

	return fi.Close()
}

// @router /certificates [get]
func (c *CertificatesController) Get() {
	c.TplName = "certificates.html"
	c.showCerts()
}

func (c *CertificatesController) showCerts() {
	path := models.GlobalCfg.OVConfigPath + "keys/index.txt"
	certs, err := lib.ReadCerts(path)
	if err != nil {
		logs.Error(err)
	}
	lib.Dump(certs)
	c.Data["certificates"] = &certs
}

// @router /certificates [post]
func (c *CertificatesController) Post() {
	c.TplName = "certificates.html"
	flash := web.NewFlash()

	cParams := NewCertParams{}
	if err := c.ParseForm(&cParams); err != nil {
		logs.Error(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	} else {
		if vMap := validateCertParams(cParams); vMap != nil {
			c.Data["validation"] = vMap
		} else {
			if err := lib.CreateCertificate(cParams.Name); err != nil {
				logs.Error(err)
				flash.Error(err.Error())
				flash.Store(&c.Controller)
			}
		}
	}
	c.showCerts()
}

func validateCertParams(cert NewCertParams) map[string]map[string]string {
	valid := validation.Validation{}
	b, err := valid.Valid(&cert)
	if err != nil {
		logs.Error(err)
		return nil
	}
	if !b {
		return lib.CreateValidationMap(valid)
	}
	return nil
}

func saveClientConfig(name string) (string, error) {
	cfg := config.New()
	cfg.ServerAddress = models.GlobalCfg.ServerAddress
	cfg.Cert = name + ".crt"
	cfg.Key = name + ".key"
	serverConfig := models.OVConfig{Profile: "default"}
	serverConfig.Read("Profile")
	cfg.Port = serverConfig.Port
	cfg.Proto = serverConfig.Proto
	cfg.Auth = serverConfig.Auth
	cfg.Cipher = serverConfig.Cipher
	cfg.Keysize = serverConfig.Keysize

	destPath := models.GlobalCfg.OVConfigPath + "keys/" + name + ".conf"
	if err := config.SaveToFile("conf/openvpn-client-config.tpl",
		cfg, destPath); err != nil {
		logs.Error(err)
		return "", err
	}

	return destPath, nil
}
