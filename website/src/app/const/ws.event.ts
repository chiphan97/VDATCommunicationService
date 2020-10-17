export enum WsEvent {
  CONNECT = 'connect',
  DISCONNECT = 'disconnect',
  CREATE_GROUP = 'group:create_group',
  UPDATE_GROUP = 'group:update_group',
  DELETE_GROUP = 'group:delete_group',
  LIST_ALL_GROUP = 'group:list_group',
  LIST_MEMBER_OF_GROUP = 'group:member:list_member',
  ADD_MEMBER_TO_GROUP = 'group:member:add_member',
  MEMBER_OUT_GROUP = 'group:member:member_out_group',
  DELETE_MEMBER_FROM_GROUP = 'group:member:delete_member',
  SEND_TEXT = 'send_text',
  SUBCRIBE_GROUP = 'subcribe_group'
}
