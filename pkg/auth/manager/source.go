// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/gitbundle/server/pkg/auth"
	"github.com/gitbundle/server/pkg/db"
	user_model "github.com/gitbundle/server/pkg/user"
)

// DeleteSource deletes a AuthSource record in DB.
func DeleteSource(source *auth.Source) error {
	count, err := db.GetEngine(db.DefaultContext).Count(&user_model.User{LoginSource: source.ID})
	if err != nil {
		return err
	} else if count > 0 {
		return auth.ErrSourceInUse{
			ID: source.ID,
		}
	}

	count, err = db.GetEngine(db.DefaultContext).Count(&user_model.ExternalLoginUser{LoginSourceID: source.ID})
	if err != nil {
		return err
	} else if count > 0 {
		return auth.ErrSourceInUse{
			ID: source.ID,
		}
	}

	if registerableSource, ok := source.Cfg.(auth.RegisterableSource); ok {
		if err := registerableSource.UnregisterSource(); err != nil {
			return err
		}
	}

	_, err = db.GetEngine(db.DefaultContext).ID(source.ID).Delete(new(auth.Source))
	return err
}

// const (
// 	localSafeIP      = "127.0.0.1"
// 	localSafeIpAlias = "::1"
// )

// func GetRealIP(r *http.Request) (string, error) {
// 	//Get IP from the X-REAL-IP header
// 	ip := r.Header.Get("X-Real-Ip")
// 	netIP := net.ParseIP(ip)
// 	if netIP != nil {
// 		return ip, nil
// 	}

// 	//Get IP from X-FORWARDED-FOR header
// 	ips := r.Header.Get("X-Forwarded-For")
// 	splitIps := strings.Split(ips, ",")
// 	for _, ip := range splitIps {
// 		netIP := net.ParseIP(ip)
// 		if netIP != nil {
// 			return ip, nil
// 		}
// 	}

// 	//Get IP from RemoteAddr
// 	ip, _, err := net.SplitHostPort(r.RemoteAddr)
// 	if err != nil {
// 		return "", err
// 	}
// 	netIP = net.ParseIP(ip)
// 	if netIP != nil {
// 		return ip, nil
// 	}
// 	return "", fmt.Errorf("No valid ip found")
// }
