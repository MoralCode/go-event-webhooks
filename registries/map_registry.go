package registries

import "github.com/MoralCode/go-event-webhooks/models"
import "errors"



type MapRegistry map[string]models.Webhooks


/* implement interface methods*/
func CreateMapRegistry() (MapRegistry) {
    /* method implementation */
    return MapRegistry{}
}


func (m_registry MapRegistry) AddToEvent(webhook models.Webhook, eventId string) (error) {
    values, ok := m_registry[eventId]   
    /* if ok is true, entry is present otherwise entry is absent*/
    if (ok) {
        if (webhook == models.Webhook{}) {
            return errors.New("Webhook cannot be empty")
        } else if (values.FindIndexOf(webhook) != -1) {
             return errors.New("Webhook already exists for this eventId")
        } else {
            m_registry[eventId] = append(values, webhook)
        }
    } else {
        // before the loop
        output := []models.Webhook{}
        output = append(output, webhook)
        m_registry[eventId] = output
    }
    return nil
}

func (m_registry MapRegistry) RemoveFromEvent(webhook models.Webhook, eventId string) (error) {
    index := m_registry[eventId].FindIndexOf(webhook)

    if index == -1 {
        return errors.New("provided webhook is not present in the registry for the given event ID")
    }

    newlist, err := m_registry[eventId].RemoveIndex(index)
    if err != nil {
        return err
    }

    m_registry[eventId] = newlist
    return nil
}

func (m_registry MapRegistry) GetHooksForEvent(eventId string) ([]models.Webhook) {
    return m_registry[eventId]
}


func (m_registry MapRegistry) Find(webhook models.Webhook) (string, int) {
    for key, _ := range m_registry {
        index := m_registry[key].FindIndexOf(webhook)

        if index != -1 {
            return key, index
        }
    }
    return "", -1
}

func (m_registry MapRegistry) FindInEvent(eventId string, webhook models.Webhook) (int, error) {
   return m_registry[eventId].FindIndexOf(webhook), nil
}

func (m_registry MapRegistry) ListEvents() ([]string) {
    keys := make([]string, 0, len(m_registry))
    for k := range m_registry {
        keys = append(keys, k)
    }
    return keys
}