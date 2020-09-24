import {WsEvent} from '../../const/ws.event';

export interface MessagePayload {
  type: WsEvent;
  groupId?: number;
  data: any;
}
