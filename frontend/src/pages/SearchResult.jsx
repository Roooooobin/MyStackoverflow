import { useParams } from "react-router";
import QuestionCard from "../components/Question/QuestionCard";
import Header from "../components/Header/Header";

function SearchResult() {
  let { q } = useParams();

  let fakeDatas = [1, 2, 3];
  const listQid = fakeDatas.map((data) => <QuestionCard qid={data} />);

  return (
    <div>
      <Header></Header>
      <h2>This is SearchResult Page!!!</h2>
      <p>The answer for "{q}" is:</p>
      <ul>{listQid}</ul>
    </div>
  );
}

export default SearchResult;
