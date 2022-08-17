export type ResponseGetUser = {
  user_id: string;
  user_name: string;
};

export type ResponseGetChannel = {
  channel_id: string;
  channel_name: string;
  channel_admin_user_id: string;
};

export type ResponseGetMessage = {
  message_id: string;
  message_send_user_id: string;
  message_send_time: string;
  message_channel_id: string;
  message_edit_flag: boolean;
  message_message: string;
};
