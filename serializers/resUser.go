package serializers

import "go-rest-service/models"

func ResUserTinySerializer(user *models.ResUser) map[string]interface{} {
	return map[string]interface{}{
		"login":          user.Login,
		"password":       user.Password,
		"password_crypt": user.PasswordCrypt,
	}
}
