package groups

import (
	"gitlab.com/vdat/mcsvc/chat/pkg/service/auth"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/userdetail"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/utils"
	"log"
	"net/http"
)

func HandLerEvent(event MessageEvent) []byte {

	switch event.Type {
	case LIST:
		user := event.UserId
		check, err := userdetail.GetUserDetailByIDService(user)
		if err != nil {
			log.Fatal(err)
			return utils.ResponseWithErrByte(http.StatusInternalServerError)
		}
		var groups []Dto
		if check.Role == userdetail.PATIENT {
			groups, err = GetGroupByPatientService(user)
			if err != nil {
				log.Fatal(err)
				return utils.ResponseWithErrByte(http.StatusInternalServerError)
			}

		} else {
			groups, err = GetGroupByDoctorService(user)
			if err != nil {
				log.Fatal(err)
				return utils.ResponseWithErrByte(http.StatusInternalServerError)
			}

		}
		event.ListGroup = groups
		return utils.ResponseWithByte(event)
	case CREATE:

		groupPayLoad := event.Data.ConvertToPayLoad()
		groupPayLoad.Users = event.Data.Users
		user := event.UserId

		check, err := userdetail.GetUserDetailByIDService(user)
		if err != nil {
			log.Fatal(err)
			return utils.ResponseWithErrByte(http.StatusInternalServerError)
		}

		if groupPayLoad.Type == ONE { //api Tạo hội thoại 1 - 1 (nhóm bí mật) ||
			_, err := GetGroupByOwnerAndUserService(groupPayLoad, user)
			if err != nil {
				log.Fatal(err)
				return utils.ResponseWithErrByte(http.StatusInternalServerError)
			}

		} else { //Tạo hội thoại n- n

			if check.Role == userdetail.PATIENT {
				return utils.ResponseWithErrByte(http.StatusForbidden)
			}

			_, err := AddGroupManyService(groupPayLoad, user)
			if err != nil {
				log.Fatal(err)
				return utils.ResponseWithErrByte(http.StatusInternalServerError)
			}

		}

		var groups []Dto
		if check.Role == userdetail.PATIENT {
			groups, err = GetGroupByPatientService(user)
			if err != nil {
				log.Fatal(err)
				return utils.ResponseWithErrByte(http.StatusInternalServerError)
			}

		} else {
			groups, err = GetGroupByDoctorService(user)
			if err != nil {
				log.Fatal(err)
				return utils.ResponseWithErrByte(http.StatusInternalServerError)
			}

		}
		event.Type = LIST
		event.ListGroup = groups
		return utils.ResponseWithByte(event)

	case UPDATE:
		owner := event.UserId
		check, err := CheckRoleOwnerInGroupService(owner, int(event.IdGroup))
		if err != nil {
			log.Fatal(err)
			return utils.ResponseWithErrByte(http.StatusInternalServerError)
		}
		if !check {
			return utils.ResponseWithErrByte(http.StatusForbidden)
		} else {
			groupPayLoad := event.Data.ConvertToPayLoad()

			newgroup, err := UpdateGroupService(groupPayLoad, int(event.IdGroup))
			if err != nil {
				log.Fatal(err)
				return utils.ResponseWithErrByte(http.StatusBadRequest)
			}
			event.IdGroup = newgroup.Id
			event.Data = newgroup.ConvertToData()
			return utils.ResponseWithByte(event)
		}
	case DELETE:

		owner := event.UserId
		check, err := CheckRoleOwnerInGroupService(owner, int(event.IdGroup))
		if !check {
			return utils.ResponseWithErrByte(http.StatusForbidden)
		} else {
			err = DeleteGroupService(int(event.IdGroup))
			if err != nil {
				log.Fatal(err)
				return utils.ResponseWithErrByte(http.StatusInternalServerError)

			}
			check, err := userdetail.GetUserDetailByIDService(owner)
			if err != nil {
				log.Fatal(err)
				return utils.ResponseWithErrByte(http.StatusInternalServerError)
			}
			var groups []Dto
			if check.Role == userdetail.PATIENT {
				groups, err = GetGroupByPatientService(owner)
				if err != nil {
					log.Fatal(err)
					return utils.ResponseWithErrByte(http.StatusInternalServerError)
				}

			} else {
				groups, err = GetGroupByDoctorService(owner)
				if err != nil {
					log.Fatal(err)
					return utils.ResponseWithErrByte(http.StatusInternalServerError)
				}

			}
			event.Type = LIST
			event.ListGroup = groups
			return utils.ResponseWithByte(event)
		}
	case LISTMEMBER:

		groupID := event.IdGroup
		users, err := GetListUserByGroupService(int(groupID))
		if err != nil {
			log.Fatal(err)
			return utils.ResponseWithErrByte(http.StatusInternalServerError)
		}
		event.ListUser = users
		return utils.ResponseWithByte(event)

	case ADDMEMBER:
		owner := auth.JWTparseOwner(event.UserId)
		check, err := CheckRoleOwnerInGroupService(owner, int(event.IdGroup))
		if !check {
			return utils.ResponseWithErrByte(http.StatusForbidden)

		} else {
			groupPayload := event.Data.ConvertToPayLoad()

			err = AddUserInGroupService(groupPayload.Users, int(event.IdGroup))
			if err != nil {
				log.Fatal(err)
				return utils.ResponseWithErrByte(http.StatusInternalServerError)
			}

			groupID := event.IdGroup
			users, err := GetListUserByGroupService(int(groupID))
			if err != nil {
				log.Fatal(err)
				return utils.ResponseWithErrByte(http.StatusInternalServerError)
			}
			event.Type = LISTMEMBER
			event.ListUser = users
			return utils.ResponseWithByte(event)

		}
	case DELETEMEMBER:

		groupID := event.IdGroup

		userid := event.Data.UserCurrent
		owner := event.UserId
		check, err := CheckRoleOwnerInGroupService(owner, int(groupID))
		if err != nil {
			log.Fatal(err)
			return utils.ResponseWithErrByte(http.StatusInternalServerError)
		}
		if !check {
			return utils.ResponseWithErrByte(http.StatusForbidden)
		} else {
			//xoa thanh vien trong nhom
			users := []string{userid}
			err := DeleteUserInGroupService(users, int(groupID))
			if err != nil {
				log.Fatal(err)
				return utils.ResponseWithErrByte(http.StatusInternalServerError)
			}

			groupID := event.IdGroup
			users2, err := GetListUserByGroupService(int(groupID))
			if err != nil {
				log.Fatal(err)
				return utils.ResponseWithErrByte(http.StatusInternalServerError)
			}
			event.Type = LISTMEMBER
			event.ListUser = users2
			return utils.ResponseWithByte(event)
		}
	case MEMBEROUT:

		groupID := event.IdGroup
		owner := event.UserId
		check, err := CheckRoleOwnerInGroupService(owner, int(groupID))
		if err != nil {
			log.Fatal(err)
			return utils.ResponseWithErrByte(http.StatusInternalServerError)
		}
		if check {
			return utils.ResponseWithErrByte(http.StatusForbidden)
		} else {
			//
			users := []string{owner}
			err := DeleteUserInGroupService(users, int(groupID))
			if err != nil {
				log.Fatal(err)
				return utils.ResponseWithErrByte(http.StatusInternalServerError)
			}

			groupID := event.IdGroup
			users2, err := GetListUserByGroupService(int(groupID))
			if err != nil {
				log.Fatal(err)
				return utils.ResponseWithErrByte(http.StatusInternalServerError)
			}
			event.Type = LISTMEMBER
			event.ListUser = users2
			return utils.ResponseWithByte(event)
		}
	default:

	}
	return nil
}
