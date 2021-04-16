package models

import "errors"

type Webhook struct {
    Url string         `json:"url"`
    HttpMethod string  `json:"httpMethod"`
}

type Webhooks []Webhook


// inspired by: https://stackoverflow.com/a/37335777/
func (list Webhooks) RemoveIndex(index int) (Webhooks, error) {
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

func (list Webhooks) FindIndexOf(webhook Webhook) (int) {
    for i, n := range list {
        if webhook == n {
            return i
        }
    }
    return -1
}