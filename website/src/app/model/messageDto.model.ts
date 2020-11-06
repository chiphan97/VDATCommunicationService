import {MessagePayload} from './payload/message.payload';
import {WsEvent} from '../const/ws.event';
import * as _ from 'lodash';

export class MessageDto {
  id: number;
  groupId: number;
  senderId: string;
  socketId: number;
  content: any;
  status: string;
  createdAt: Date;
  type: WsEvent;

  constructor(id: number, groupId: number, senderId: string, socketId: number,
              content: any, status: string, createdAt: Date = new Date()) {
    this.id = id;
    this.groupId = groupId;
    this.senderId = senderId;
    this.socketId = socketId;
    this.content = content;
    this.status = status;
    this.createdAt = createdAt;
  }

  public static fromJson(data: any): MessageDto {
    if (!!data) {
      const messageData = _.get(data, 'data', null);

      if (!!messageData) {
        const messageDto = new MessageDto(
          _.get(messageData, 'id', -1),
          _.get(messageData, 'groupId', -1),
          _.get(messageData, 'Sender', ''),
          _.get(messageData, 'socketId', -1),
          _.get(messageData, 'body', ''),
          _.get(messageData, 'Status', '')
        );

        messageDto.type = _.get(data, 'type', WsEvent.MESSAGE);

        if (!!!messageDto.senderId) {
          messageDto.senderId = _.get(data, 'Client', '');
        }

        return messageDto;
      }
    }

    return null;
  }
}
