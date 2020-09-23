import {GroupType} from '../../const/group-type.const';

export class GroupPayload {
  nameGroup: string;
  type: GroupType;
  private: boolean;
  users: Array<string>;
  description: string;
}
