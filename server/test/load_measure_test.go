package test

import (
	"bytes"
	"cow_backend_mobile/models"
	"cow_backend_mobile/routes"
	"cow_backend_mobile/routes/load"
	"cow_backend_mobile/routes/user"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	r.Use(cors.New(cors.Config{
		AllowOrigins:           []string{"*"},
		AllowMethods:           []string{"PUT", "PATCH", "POST", "OPTIONS", "DELETE", "GET", "HEAD"},
		AllowHeaders:           []string{"content-type", "Authorization", "authorization"},
		ExposeHeaders:          []string{"Content-Length"},
		AllowCredentials:       true,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		MaxAge:                 12 * time.Hour,
	}))

	apiGroup := r.Group("/api/mobile")
	routes.WriteRoutes(apiGroup, &load.Load{}, &user.User{})
	return r
}

func getAccessToken(router *gin.Engine, t *testing.T) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("12345"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	db := models.GetDatabase()
	newUser := models.User{
		Email:                 "test@test.com",
		Password:              hashedPassword,
		NameSurnamePatronymic: "Pupupupu",
	}
	db.FirstOrCreate(&newUser)
	db.Save(&newUser)

	w := httptest.NewRecorder()
	bodyLogin := user.LoginData{
		Email:    "test@test.com",
		Password: "12345",
	}
	out, err := json.Marshal(bodyLogin)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/api/mobile/user/login", bytes.NewBuffer(out))
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(w, req)
	resp := map[string]any{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	access := resp["access"].(string)
	assert.Equal(t, 200, w.Code)
	return access
}

func TestUnauthorizedRestrict(t *testing.T) {
	router := setupRouter()
	models.MockDatabase()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/mobile/load/measuredData", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
}

func TestEmptyBodyUnprocessableEntity(t *testing.T) {
	router := setupRouter()
	models.MockDatabase()

	access := getAccessToken(router, t)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/mobile/load/measuredData", nil)
	req.Header.Set("Authorization", access)
	router.ServeHTTP(w, req)
	assert.Equal(t, 422, w.Code)
}

func TestValidBodyField(t *testing.T) {
	router := setupRouter()
	models.MockDatabase()

	access := getAccessToken(router, t)
	body := `{
  "additionalInfo": {
    "additionalProperty1Measure": "string",
    "additionalProperty1Name": "string",
    "additionalProperty1Value": "string",
    "additionalProperty2Measure": "string",
    "additionalProperty2Name": "string",
    "additionalProperty2Value": "string",
    "calvingDate": "2001-03-23",
    "firstMilkingDate": "2001-03-23",
    "lactationNumber": 0
  },
  "cow": {
    "birthDate": "2001-03-23",
    "bloody": 1,
    "breedName": "Порода",
    "holdingInn": "string",
    "holdingName": "string",
    "hozName": "string",
    "inventoryNumber": "1213321",
    "name": "Дима",
    "rshnnumber": "1323323232",
    "selecsNumber": 98989
  },
  "downSides": null,
  "exterior": {
    "assessmentDate": "2001-03-23",
    "backBoneQuality": 9,
    "backNippleDiameter": 9,
    "backNippleLocationBackView": 9,
    "backUdderSegmentsLocationHeight": 9,
    "backUdderSegmentsWidth": 9,
    "body": 100,
    "bodyDepth": 9,
    "bodyType": 9,
    "category": "string",
    "centralLigamentDepth": 9,
    "chestWidth": 9,
    "deception": 9,
    "fatness": 9,
    "forelegWalk": 9,
    "frontNippleDiameter": 9,
    "frontNippleLength": 9,
    "frontNippleLocationBackView": 9,
    "frontUdderSegmentsLocation": 9,
    "harmonyOfMovement": 9,
    "hindLegWalkBackView": 9,
    "hindLegWalkSideView": 9,
    "hoofAngle": 9,
    "limbs": 100,
    "milkType": 100,
    "rating": 100,
    "ribsAngle": 9,
    "sacrum": 100,
    "sacrumAngle": 9,
    "sacrumHeight": 9,
    "sacrumLength": 9,
    "sacrumWidth": 9,
    "topLine": 9,
    "udder": 100,
    "udderBalance": 9,
    "udderDepth": 9,
    "udderVeins": 9
  },
  "measures": {
    "backNippleDiameter": 0,
    "backUdderSegmentsLocationHeight": 0,
    "backUdderSegmentsWidth": 0,
    "centralLigamentDepth": 0,
    "chestWidth": 0,
    "frontNippleDiameter": 0,
    "frontNippleLength": 0,
    "frontUdderSegmentsLocation": 0,
    "hindLegWalkSideView": 0,
    "hoofAngle": 0,
    "sacrumAngle": 0,
    "sacrumHeight": 0,
    "sacrumWidth": 0,
    "udderBalance": 0,
    "udderDepth": 0
  },
  "ratings": {
    "automaticWithDownsidesBody": 0,
    "automaticWithDownsidesLimbs": 0,
    "automaticWithDownsidesMilkType": 0,
    "automaticWithDownsidesSacrum": 0,
    "automaticWithDownsidesUdder": 0,
    "automaticWithoutDownsidesBody": 0,
    "automaticWithoutDownsidesLimbs": 0,
    "automaticWithoutDownsidesMilkType": 0,
    "automaticWithoutDownsidesSacrum": 0,
    "automaticWithoutDownsidesUdder": 0,
    "userDefinedWithDownsidesBody": 0,
    "userDefinedWithDownsidesLimbs": 0,
    "userDefinedWithDownsidesMilkType": 0,
    "userDefinedWithDownsidesSacrum": 0,
    "userDefinedWithDownsidesUdder": 0,
    "userDefinedWithoutDownsidesBody": 0,
    "userDefinedWithoutDownsidesLimbs": 0,
    "userDefinedWithoutDownsidesMilkType": 0,
    "userDefinedWithoutDownsidesSacrum": 0,
    "userDefinedWithoutDownsidesUdder": 0
  },
  "weights": {
    "automaticBody": 0,
    "automaticLimbs": 0,
    "automaticMilkType": 0,
    "automaticSacrum": 0,
    "automaticUdder": 0,
    "backBoneQuality": 0,
    "backNippleDiameter": 0,
    "backNippleLocationBackView": 0,
    "backUdderSegmentsLocationHeight": 0,
    "backUdderSegmentsWidth": 0,
    "bodyDepth": 0,
    "bodyType": 0,
    "centralLigamentDepth": 0,
    "chestWidth": 0,
    "deception": 0,
    "fatness": 0,
    "forelegWalk": 0,
    "frontNippleDiameter": 0,
    "frontNippleLength": 0,
    "frontNippleLocationBackView": 0,
    "frontUdderSegmentsLocation": 0,
    "harmonyOfMovement": 0,
    "hindLegWalkBackView": 0,
    "hindLegWalkSideView": 0,
    "hoofAngle": 0,
    "ribsAngle": 0,
    "sacrumAngle": 0,
    "sacrumHeight": 0,
    "sacrumLength": 0,
    "sacrumWidth": 0,
    "topLine": 0,
    "udderBalance": 0,
    "udderDepth": 0,
    "udderVeins": 0,
    "userDefinedBody": 0,
    "userDefinedLimbs": 0,
    "userDefinedMilkType": 0,
    "userDefinedSacrum": 0,
    "userDefinedUdder": 0
  }
}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/mobile/load/measuredData", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Authorization", access)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	db := models.GetDatabase()
	extCount := int64(0)
	cowCount := int64(0)
	db.Model(&models.Exterior{}).Count(&extCount)
	db.Model(&models.Cow{}).Count(&cowCount)
	assert.Equal(t, int64(1), cowCount)
	assert.Equal(t, int64(1), extCount)
}

func TestLoadingTwice(t *testing.T) {
	router := setupRouter()
	models.MockDatabase()

	access := getAccessToken(router, t)
	body := `{
  "additionalInfo": {
    "additionalProperty1Measure": "string",
    "additionalProperty1Name": "string",
    "additionalProperty1Value": "string",
    "additionalProperty2Measure": "string",
    "additionalProperty2Name": "string",
    "additionalProperty2Value": "string",
    "calvingDate": "2001-03-23",
    "firstMilkingDate": "2001-03-23",
    "lactationNumber": 0
  },
  "cow": {
    "birthDate": "2001-03-23",
    "bloody": 1,
    "breedName": "Порода",
    "holdingInn": "string",
    "holdingName": "string",
    "hozName": "string",
    "inventoryNumber": "1213321",
    "name": "Дима",
    "rshnnumber": "1323323232",
    "selecsNumber": 98989
  },
  "downSides": null,
  "exterior": {
    "assessmentDate": "2001-03-23",
    "backBoneQuality": 9,
    "backNippleDiameter": 9,
    "backNippleLocationBackView": 9,
    "backUdderSegmentsLocationHeight": 9,
    "backUdderSegmentsWidth": 9,
    "body": 100,
    "bodyDepth": 9,
    "bodyType": 9,
    "category": "string",
    "centralLigamentDepth": 9,
    "chestWidth": 9,
    "deception": 9,
    "fatness": 9,
    "forelegWalk": 9,
    "frontNippleDiameter": 9,
    "frontNippleLength": 9,
    "frontNippleLocationBackView": 9,
    "frontUdderSegmentsLocation": 9,
    "harmonyOfMovement": 9,
    "hindLegWalkBackView": 9,
    "hindLegWalkSideView": 9,
    "hoofAngle": 9,
    "limbs": 100,
    "milkType": 100,
    "rating": 100,
    "ribsAngle": 9,
    "sacrum": 100,
    "sacrumAngle": 9,
    "sacrumHeight": 9,
    "sacrumLength": 9,
    "sacrumWidth": 9,
    "topLine": 9,
    "udder": 100,
    "udderBalance": 9,
    "udderDepth": 9,
    "udderVeins": 9
  },
  "measures": {
    "backNippleDiameter": 0,
    "backUdderSegmentsLocationHeight": 0,
    "backUdderSegmentsWidth": 0,
    "centralLigamentDepth": 0,
    "chestWidth": 0,
    "frontNippleDiameter": 0,
    "frontNippleLength": 0,
    "frontUdderSegmentsLocation": 0,
    "hindLegWalkSideView": 0,
    "hoofAngle": 0,
    "sacrumAngle": 0,
    "sacrumHeight": 0,
    "sacrumWidth": 0,
    "udderBalance": 0,
    "udderDepth": 0
  },
  "ratings": {
    "automaticWithDownsidesBody": 0,
    "automaticWithDownsidesLimbs": 0,
    "automaticWithDownsidesMilkType": 0,
    "automaticWithDownsidesSacrum": 0,
    "automaticWithDownsidesUdder": 0,
    "automaticWithoutDownsidesBody": 0,
    "automaticWithoutDownsidesLimbs": 0,
    "automaticWithoutDownsidesMilkType": 0,
    "automaticWithoutDownsidesSacrum": 0,
    "automaticWithoutDownsidesUdder": 0,
    "userDefinedWithDownsidesBody": 0,
    "userDefinedWithDownsidesLimbs": 0,
    "userDefinedWithDownsidesMilkType": 0,
    "userDefinedWithDownsidesSacrum": 0,
    "userDefinedWithDownsidesUdder": 0,
    "userDefinedWithoutDownsidesBody": 0,
    "userDefinedWithoutDownsidesLimbs": 0,
    "userDefinedWithoutDownsidesMilkType": 0,
    "userDefinedWithoutDownsidesSacrum": 0,
    "userDefinedWithoutDownsidesUdder": 0
  },
  "weights": {
    "automaticBody": 0,
    "automaticLimbs": 0,
    "automaticMilkType": 0,
    "automaticSacrum": 0,
    "automaticUdder": 0,
    "backBoneQuality": 0,
    "backNippleDiameter": 0,
    "backNippleLocationBackView": 0,
    "backUdderSegmentsLocationHeight": 0,
    "backUdderSegmentsWidth": 0,
    "bodyDepth": 0,
    "bodyType": 0,
    "centralLigamentDepth": 0,
    "chestWidth": 0,
    "deception": 0,
    "fatness": 0,
    "forelegWalk": 0,
    "frontNippleDiameter": 0,
    "frontNippleLength": 0,
    "frontNippleLocationBackView": 0,
    "frontUdderSegmentsLocation": 0,
    "harmonyOfMovement": 0,
    "hindLegWalkBackView": 0,
    "hindLegWalkSideView": 0,
    "hoofAngle": 0,
    "ribsAngle": 0,
    "sacrumAngle": 0,
    "sacrumHeight": 0,
    "sacrumLength": 0,
    "sacrumWidth": 0,
    "topLine": 0,
    "udderBalance": 0,
    "udderDepth": 0,
    "udderVeins": 0,
    "userDefinedBody": 0,
    "userDefinedLimbs": 0,
    "userDefinedMilkType": 0,
    "userDefinedSacrum": 0,
    "userDefinedUdder": 0
  }
}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/mobile/load/measuredData", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Authorization", access)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/mobile/load/measuredData", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Authorization", access)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	db := models.GetDatabase()
	extCount := int64(0)
	cowCount := int64(0)
	db.Model(&models.Exterior{}).Count(&extCount)
	db.Model(&models.Cow{}).Count(&cowCount)
	assert.Equal(t, int64(1), cowCount)
	assert.Equal(t, int64(2), extCount)
}

//
