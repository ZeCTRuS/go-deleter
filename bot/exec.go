/*
 * Copyleft (ↄ) 2020, Geosonic
 */

package bot

import (
	"fmt"

	"github.com/SevereCloud/vksdk/api"
)

/*

Execute функции

*/

// За рефакторинг execute кода спасибо https://vk.com/notqb
/*
// Количество которое необходимо удалить
// Эти переменные устанавливаются скриптом!!!
var delete_count = %v, peer_id = %v; // int

// Переменная отсортированных сообщений
var message_ids = [];

// Получаем сообщения
var messages = API.messages.getHistory({peer_id: peer_id, count: 200}).items + API.messages.getHistory({peer_id: peer_id, count: 200, offset: 200}).items;
// Получаем ID аккаунта
var self_id = API.users.get()[0].id;

while (messages.length > 0 && message_ids.length < delete_count) {
// Переменная сообщения
var message = messages.shift();

// Находим свои сообщения
if (message.from_id == self_id) {
message_ids.push(message.id);
}
}
*/
const code = `var delete_count=%v,peer_id=%v;var message_ids=[];var messages=API.messages.getHistory({peer_id:peer_id,count:200}).items+API.messages.getHistory({peer_id:peer_id,count:200,offset:200}).items;var self_id=API.users.get()[0].id;while(messages.length>0&&message_ids.length<delete_count){var message=messages.shift();if(message.from_id==self_id){message_ids.push(message.id);}}`

func DeleteExec(vk *api.VK, toDeleteCount, peerID int) error {
	code := // Возвращаем результат удаления сообщений
		fmt.Sprintf(code+`return API.messages.delete({message_ids: message_ids, delete_for_all: 1});`, toDeleteCount, peerID)

	err := vk.Execute(code, nil)

	return err
}

func GetMessages(vk *api.VK, toDeleteCount, peerID int) ([]int, error) {
	code := // Возвращаем найденные сообщения
		fmt.Sprintf(code+`return {messages_ids: message_ids};`, toDeleteCount, peerID)

	var resp struct {
		Messages []int `json:"messages_ids"`
	}

	err := vk.Execute(code, &resp)

	return resp.Messages, err
}
