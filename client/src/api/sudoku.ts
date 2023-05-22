import { useEffect, useState } from "react";

type Board = number[][];

interface GetNewGameResponse {
  board: Board;
}

const baseUrl = "http://localhost:8001";

const getNewGame = async (): Promise<GetNewGameResponse> => {
  const response = await fetch(`${baseUrl}/v1/sudoku/`, {
    method: "GET",
  });

  const json = await response.json();

  return {
    board: json.board,
  };
};

export const useNewGame = () => {
  const [board, setBoard] = useState<Board>();
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [isError, setIsError] = useState<boolean>(false);

  useEffect(() => {
    const api = async () => {
      setIsLoading(true);
      try {
        const game = await getNewGame();
        setBoard(game.board);
      } catch (error) {
        console.error(error);
        setIsError(true);
      } finally {
        setIsLoading(false);
      }
    };
    api();
  }, []);

  return { data: board, isLoading, isError };
};

interface PostSolveGameResponse {
  board: Board;
  isSolved: boolean;
}

export const postSolveGame = async (
  board: Board
): Promise<PostSolveGameResponse> => {
  const response = await fetch(`${baseUrl}/v1/sudoku/solve`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ board }),
  });

  const json = await response.json();

  console.log(json.isSolved);

  return {
    board: json.board,
    isSolved: json.isSolved,
  };
};
