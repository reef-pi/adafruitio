## go client for adafruit.io

This library provides a go client for [v2 adafruit.io rest API](https://io.adafruit.com/api/docs/#!/v2). Adafruit.io
provides an easy way to log and chart data originting from IoT devices. It can also send webhooks or emails (called
as triggers) based on data.

Adafruit.com provides a comprehensive [tutorial](https://learn.adafruit.com/adafruit-io/overview) for adafruit.io services.

## Usage

```go
import(
 "github.com/reef-pi/adafruitio"
)

func demo(){
  client := adafruitio.NewClient("<adafruit.io token>")
  opt := adafruitio.ListActivitiesOptions{
    Limit: 100,
  }
  activities, err := client.ListActivities("example_user", opt)
  if err != nil {
    // deal with it
    return
  }
  for _, a := range activities {
    fmt.Println(a.Model, a.Action, a.CreatedAt) 
  }

}
 
```


## License

Copyright:: Copyright (c) 2017 Ranjib Dey.
License:: Apache License, Version 2.0

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
