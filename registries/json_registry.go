package registries

import "errors"
import "github.com/MoralCode/go-event-webhooks/models"
import "encoding/json"
import "log"
import "io/ioutil"
import "os"

type JSONRegistry struct {
    Registry MapRegistry;
    FilePath string;
}


/* implement interface methods*/
func CreateNewJSONRegistry(filepath string) (JSONRegistry) {
    /* method implementation */
    return JSONRegistry{MapRegistry{}, filepath}
}

func CreateJSONRegistryFromMapRegistry(mapregistry MapRegistry) (JSONRegistry) {
    /* method implementation */
    return JSONRegistry{mapregistry, ""}
}

func CreateJSONRegistryFromJSONFile(filename string) (JSONRegistry) {
    map_registry := readFromDisk(filename)

    registry := CreateJSONRegistryFromMapRegistry(map_registry)
    registry.FilePath = filename
    return registry
}


func (j_registry JSONRegistry) AddToEvent(webhook models.Webhook, eventId string) (error) {
    err := j_registry.Registry.AddToEvent(webhook, eventId)
    if err != nil {
        return err
    }
    err = j_registry.writeToDisk()
    if err != nil {
        return err
    }
    return nil
}

func (j_registry JSONRegistry) RemoveFromEvent(webhook models.Webhook, eventId string) (error) {
    err := j_registry.Registry.RemoveFromEvent(webhook, eventId)
    if err != nil {
        return err
    }
    err = j_registry.writeToDisk()
    if err != nil {
        return err
    }
    return nil
}

func (j_registry JSONRegistry) GetHooksForEvent(eventId string) ([]models.Webhook) {
    return j_registry.Registry.GetHooksForEvent(eventId)
}

func (j_registry JSONRegistry) Find(webhook models.Webhook) (string, int) {
     return j_registry.Registry.Find( webhook)
}

func (j_registry JSONRegistry) FindInEvent(eventId string, webhook models.Webhook) (int, error) {
    return j_registry.Registry.FindInEvent(eventId, webhook)
}

func (j_registry JSONRegistry) ListEvents() ([]string) {
   return j_registry.Registry.ListEvents()
}

func (j_registry JSONRegistry) writeToDisk() (error) {
    if j_registry.FilePath == "" {
        return errors.New("Filepath of JSONRegistry must not be empty")
    }

    file, err := json.Marshal(j_registry.Registry)
    if err != nil {
        return err
    }

    err = ioutil.WriteFile(j_registry.FilePath, file, 0644)
    if err != nil {
        return err
    }
    return nil
}

func readFromDisk(filename string) (MapRegistry) {

    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    data, err := ioutil.ReadAll(file)

    if err != nil {
        log.Fatal(err)
    }

    var result MapRegistry

    jsonErr := json.Unmarshal(data, &result)

    if jsonErr != nil {
        log.Fatal(jsonErr)
    }

    return result

}

