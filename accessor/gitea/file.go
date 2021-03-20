// Copyright 2020 Steve Jefferson. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package gitea

import (
	"io"
	"os"
	"path/filepath"
	"github.com/pkg/errors"
	"github.com/stevejefferson/trac2gitea/log"
)

func copyFile(externalFilePath string, giteaPath string) error {
	_, err := os.Stat(externalFilePath)
	if os.IsNotExist(err) {
		log.Warn("cannot copy non-existant attachment file: \"%s\"", externalFilePath)
		return nil
	}

	in, err := os.Open(externalFilePath)
	if err != nil {
		err = errors.Wrapf(err, "opening file %s", externalFilePath)
		return err
	}
	defer in.Close()

	if err := os.MkdirAll(filepath.Dir(giteaPath), 0770); err != nil {
        	return err
    	}

	out, err := os.Create(giteaPath)
	if err != nil {
		err = errors.Wrapf(err, "creating file %s", giteaPath)
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		err = errors.Wrapf(err, "copying %s to %s", externalFilePath, giteaPath)
		return err
	}

	err = out.Close()
	if err != nil {
		err = errors.Wrapf(err, "closing %s", giteaPath)
		return err
	}

	log.Debug("copied file %s to %s", externalFilePath, giteaPath)
	return nil
}

func deleteFile(path string) error {
	return os.Remove(path)
}
