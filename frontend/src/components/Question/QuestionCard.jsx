import { Link } from "react-router-dom";

function QuestionCard({ qid }) {
  return (
    <div className="questionCard">
      <Link to={`/question/${qid}`}>
        <label>The uid is: {qid}</label>
      </Link>
    </div>
  );
}

export default QuestionCard;
