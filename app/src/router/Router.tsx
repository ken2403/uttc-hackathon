import { FC } from "react";
import { Switch, Route } from "react-router-dom";

import { Login } from "../components/pages/Login";
import { ChannelIndex } from "../components/pages/ChannelIndex";
import { Chat } from "../components/pages/Chat";
import { Page404 } from "../components/pages/Page404";

export const Router: FC = () => {
  return (
    <Switch>
      <Route exact path="/">
        <Login />
      </Route>
      <Route exact path="/channelindex/*">
        <ChannelIndex />
      </Route>
      <Route exact path="/chat/*">
        <Chat channel={null} />
      </Route>
      <Route path="*">
        <Page404 />
      </Route>
    </Switch>
  );
};
