package registries

import "github.com/MoralCode/go-event-webhooks/models"
import "errors"

type Webhooks []models.Webhook

type MapRegistry map[string]Webhooks


/* implement interface methods*/
func CreateMapRegistry() (MapRegistry) {
    /* method implementation */
    return MapRegistry{}
}


func (m_registry MapRegistry) AddToEvent(webhook models.Webhook, eventId string) {
    values, ok := m_registry[eventId]   
    /* if ok is true, entry is present otherwise entry is absent*/
    if (ok) {
        //TODO: err if exists already?
        if (webhook != models.Webhook{} && values.findIndexOf(webhook) == -1) {
                m_registry[eventId] = append(values, webhook)
        }
    } else {
        // before the loop
        output := []models.Webhook{}
        output = append(output, webhook)
        m_registry[eventId] = output
    }
}

func (m_registry MapRegistry) RemoveFromEvent(webhook models.Webhook, eventId string) (error) {
        index := m_registry[eventId].findIndexOf(webhook)

        if index == -1 {
            return errors.New("provided webhook is not present in the registry for the given event ID")
        }

        newlist, err := m_registry[eventId].removeIndex(index)
        if err != nil {
            return err
        }

        m_registry[eventId] = newlist
        return nil
    }

// inspired by: https://stackoverflow.com/a/37335777/
func (list Webhooks) removeIndex(index int) (Webhooks, error) {
    if index < 0 {
        return list, errors.New("negative indices are not allowed")
    }
    
    if (index >= len(list)) {
        return list, errors.New("provided index exceeds the bounds of the provided list ")
    }

    if (index == len(list)-1) {
        //if the item to remove is the last one, just return all but the last item
        return list[:len(list)-1], nil
    }

    //otherwise replace the item to be removed with the last item

    list[index] = list[len(list)-1]
    // We do not need to put s[i] at the end, as it will be discarded anyway
    return list[:len(list)-1], nil
}

func (list Webhooks) findIndexOf(webhook models.Webhook) (int) {
    for i, n := range list {
        if webhook == n {
            return i
        }
    }
    return -1
}

func (m_registry MapRegistry) Find(webhook models.Webhook) (string, int) {
    for key, _ := range m_registry {
        index := m_registry[key].findIndexOf(webhook)

        if index != -1 {
            return key, index
        }
    }
    return "", -1
}

func (m_registry MapRegistry) FindInEvent(eventId string, webhook models.Webhook) (int, error) {
   return m_registry[eventId].findIndexOf(webhook), nil
}

func (m_registry MapRegistry) ListEvents() ([]string) {
    keys := make([]string, 0, len(m_registry))
    for k := range m_registry {
        keys = append(keys, k)
    }
    return keys
}