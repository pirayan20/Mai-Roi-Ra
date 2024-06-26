"use server";
import { getServerSession } from "next-auth";
import getProblem from "@/libs/getProblem";
import ReplyForm from "@/components/ReplyForm";
import { authOptions } from "@/app/api/auth/[...nextauth]/auth";

interface Props {
  params: { problem_id: string };
}

export default async function ReplyProblemPage({ params }: Props) {
  // This is for ADMIN only

  const session = await getServerSession(authOptions);
  if (!session || !session.user || !session.user.token) return null;
  const problem = await getProblem(params.problem_id, session.user.token);
  return (
    <main className="h-auto w-full text-black">
      <ReplyForm
        problem={problem.problem}
        description={problem.description}
        problem_id={params.problem_id}
        token={session.user.token}
      ></ReplyForm>
    </main>
  );
}
