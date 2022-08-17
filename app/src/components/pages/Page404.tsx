import { Link } from "react-router-dom";
import { BaseContainer } from "../atoms/container/BaseContainer";

export const Page404 = () => {
  return (
    <BaseContainer>
      <h1>ページが見つかりません</h1>
      <Link to="/">Topに戻る</Link>
    </BaseContainer>
  );
};
