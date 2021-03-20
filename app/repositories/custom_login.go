package repositories

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_core "github.com/kiketordera/advanced-performance/app/core"
	_domain "github.com/kiketordera/advanced-performance/app/domain"
)

// LoginRepository implements models.ILogin with Custom Login
type LoginRepository struct {
}

// InitDatabase creates the instance of the BoltHold Database
func GetCustomLogin() _domain.ILogin {
	return LoginRepository{}
}

// TokenClaims is the token that we store in the cookies
type TokenClaims struct {
	Username string `json:"usr,omitempty"`
	Role     string `json:"rol,omitempty"`
	jwt.StandardClaims
}

// This middleware ensures that a request will be aborted with an error if the user is not logged in
// This method checks if the user if logged in
func (r LoginRepository) CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// If there's an error or if the token is empty
		// the user is not logged in
		cookie, err := c.Cookie(_core.KNameProyect)
		if err != nil {
			r.SetSession(c, "", "")
			return
		}
		token, err := jwt.ParseWithClaims(cookie, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Unexpected signing method: " + token.Header["alg"].(string))
			}

			// Return signing key to check the token
			return _core.TokenSigningKey, nil
		})
		if err != nil {
			print(err)
		}
		if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
			c.Set("username", claims.Username)
			c.Set("role", claims.Role)
			if claims.Username != "" {
				c.Set("isLoggedIn", true)
			} else {
				c.Set("isLoggedIn", false)
			}
		} else {
			c.Set("username", "")
			c.Set("role", "")
			c.Set("isLoggedIn", false)
			fmt.Println("Redirecting to /login from check token")
			c.Redirect(http.StatusFound, "/login")
		}
	}
}

// Set the session in the cookie
func (r LoginRepository) SetSession(c *gin.Context, username string, role _domain.UserRole) {
	// Define token claims
	claims := TokenClaims{
		username,
		string(role),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(_core.SessionTime),
			Issuer:    _core.KNameProyect,
		},
	}
	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(_core.TokenSigningKey)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.SetCookie(_core.KNameProyect, tokenString, _core.SessionTime, "", "", false, true)
}

// Ensures that the user is logged in
func (r LoginRepository) EnsureLoggedIn(c *gin.Context) {
	u, e := c.Get("username")
	if e {
		username := u.(string)
		if username == "" {
			fmt.Println("Redirecting to /login from ensureLoggedIn")
			c.Redirect(http.StatusFound, "/login")
		}
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

func (r LoginRepository) IsLoggedIn(c *gin.Context) bool {
	u, e := c.Get("username")
	if e {
		username := u.(string)
		return username != ""
	}
	return false
}

// Ensures that the user is NOT logged in
func (r LoginRepository) EnsureNotLoggedIn(c *gin.Context) {
	u, e := c.Get("username")
	// If the User is logged in, we send him to the Dashboard
	if e {
		username := u.(string)
		if username != "" {
			//Redirection when is logged
			fmt.Println("Redirecting to /dashboard from ensureNotLoggedIn")
			c.Redirect(http.StatusFound, "/dashboard")
		}
	}
}

// Logs out of the app for the user
func (r LoginRepository) Logout(c *gin.Context) {
	// Clear the cookie
	c.SetCookie(_core.KNameProyect, "", -1, "", "", false, true)
	c.Set("isLoggedIn", false)
	// Redirect to the home page
	fmt.Println("Redirecting to /login from logout")
	c.Redirect(http.StatusFound, "/")
}

// Ensures that the user is an Admin
func (r LoginRepository) EnsureIsAdmin(c *gin.Context) {
	a, e := c.Get("role")
	// If the User is logged in, we send him to the Dashboard
	if e {
		role := a.(string)
		if role != string(_domain.Admin) {
			//Redirection when is NOT admin
			fmt.Println("Redirecting to /dashboard from ensureIsAdmin")
			c.Redirect(http.StatusFound, "/dashboard")
		}
	}
}
