import axios, { AxiosResponse, AxiosError } from "axios";
import { FC, useEffect, useState } from "react";
import { useLocation, useHistory } from "react-router-dom";
import { URLs } from "../../api/urls";
import { Channel } from "../../types/channel";
import { User } from "../../types/user";
import { ResponseGetChannel, ResponseGetUser } from "../../types/response";
import { PrimaryButton } from "../atoms/button/PrimaryButton";
import { SecondaryButton } from "../atoms/button/SecondaryButton";
import { BaseContainer } from "../atoms/container/BaseContainer";
import { Page404 } from "./Page404";

type State =
  | undefined
  | {
      userName: string;
    };

export const ChannelIndex: FC = () => {
  const { state } = useLocation<State>();
  const [channels, setChannels] = useState<Channel[]>([]);
  const [user, setUser] = useState<User>();
  const history = useHistory();

  useEffect(() => {
    if (typeof state !== "undefined") {
      let url = URLs.getChannels;
      url += state?.userName;
      const userChannels: Channel[] = [];
      axios
        .get(url)
        .then((res: AxiosResponse) => {
          if (res.data.length !== 0) {
            const responseChannels: ResponseGetChannel[] = res.data;
            responseChannels.map((channel: ResponseGetChannel) => {
              userChannels.push({
                channelId: channel.channel_id,
                channelName: channel.channel_name,
                channelAdminUserId: channel.channel_admin_user_id
              });
            });
            setChannels(userChannels);
          }
        })
        .catch((error: AxiosError<{ error: string }>) => {
          console.log(error);
        });

      let url2 = URLs.getUser;
      url2 += state?.userName;
      axios
        .get(url2)
        .then((res: AxiosResponse) => {
          if (res.data.length !== 0) {
            const responseUsers: ResponseGetUser[] = res.data;
            const user: User = {
              userId: responseUsers[0].user_id,
              userName: responseUsers[0].user_name
            };
            setUser(user);
          }
        })
        .catch((error: AxiosError<{ error: string }>) => {
          console.log(error);
        });
    }
  }, [state]);

  const onClick = (channel: Channel) => {
    history.push(`/chat/${channel.channelName}}`, {
      channel: channel,
      user: user
    });
  };
  const onClickBack = () => {
    history.goBack();
  };

  // rendering
  if (typeof state === "undefined") {
    return <Page404 />;
  } else if (channels.length === 0) {
    return (
      <BaseContainer>
        <h1> Hello {state?.userName}</h1>
        <h2>あなたの所属しているチャンネルはありません</h2>
      </BaseContainer>
    );
  } else {
    return (
      <>
        <BaseContainer>
          <h1> Hello {state?.userName}</h1>
          <h2>あなたの所属しているチャンネル一覧</h2>
          {channels.map((channel) => (
            <BaseContainer key={channel.channelId}>
              <PrimaryButton onClick={() => onClick(channel)}>
                # {channel.channelName}
              </PrimaryButton>
            </BaseContainer>
          ))}
        </BaseContainer>
        <SecondaryButton onClick={onClickBack}>
          ユーザー選択に戻る
        </SecondaryButton>
      </>
    );
  }
};
