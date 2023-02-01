// Copyright © 2023 Weald Technology Trading.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validatorexit

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

//nolint:unparam
func (c *command) output(_ context.Context) (string, error) {
	if c.quiet {
		return "", nil
	}

	if c.prepareOffline {
		return fmt.Sprintf("%s generated", offlinePreparationFilename), nil
	}

	if c.json || c.offline {
		data, err := json.Marshal(c.signedOperation)
		if err != nil {
			return "", errors.Wrap(err, "failed to marshal signed operation")
		}
		if c.json {
			return string(data), nil
		}
		if err := os.WriteFile(exitOperationFilename, data, 0o600); err != nil {
			return "", errors.Wrap(err, fmt.Sprintf("failed to write %s", exitOperationFilename))
		}
		return "", nil
	}

	return "", nil
}
