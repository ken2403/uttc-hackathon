import { FC, useEffect, useState } from "react";
import axios, { AxiosResponse, AxiosError } from "axios";
import { FaUser } from "react-icons/fa";

import { URLs } from "../../api/urls";
import { ResponseGetUser } from "../../types/response";
import { CircleContainer } from "../atoms/container/CircleContainer";

type Props = {
  userId: string;
};

export const ShowUserName: FC<Props> = (props) => {
  const { userId } = props;
  const [userName, setUserName] = useState<string>("");

  useEffect(() => {
    let url = URLs.getUserByID;
    url += userId;
    axios
      .get(url)
      .then((res: AxiosResponse) => {
        if (res.data.length !== 0) {
          const responseUser: ResponseGetUser[] = res.data;
          const newUserName = responseUser[0].user_name;
          setUserName(newUserName);
        }
      })
      .catch((error: AxiosError<{ error: string }>) => {
        console.log(error);
      });
  }, [userId]);
  return (
    <CircleContainer>
      <FaUser />
      <br />
      {userName}
    </CircleContainer>
  );
};
