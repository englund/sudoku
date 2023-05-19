"use client";

import { postSolveGame, useNewGame } from "@/api/sudoku";
import { FC, useEffect, useState } from "react";

const SudokuGame: FC = () => {
  const { data, isLoading, isError } = useNewGame();
  const [board, setBoard] = useState(data);
  const [isSolvable, setIsSolvable] = useState<boolean>();

  useEffect(() => {
    if (data == null) return;
    setBoard(data);
  }, [data]);

  if (isLoading) return <span>loading...</span>;

  if (isError)
    return <span>something went horrible wrong, maybe try again</span>;

  if (board == null) return null;

  const updateValue = (x: number, y: number, value: string) => {
    const newBoard = [...board];
    newBoard[x][y] = +value; // TODO: validate number
    setBoard(newBoard);
  };

  const solveGame = async () => {
    const response = await postSolveGame(board);
    setIsSolvable(response.isSolved);
  };

  return (
    <div>
      <div className="[&>*:nth-child(3n+4)]:border-t-4">
        {board.map((row, x) => {
          return (
            <div
              key={x}
              className="grid grid-cols-9 [&>*:nth-child(3n+4)]:border-l-4"
            >
              {row.map((value, y) => {
                return (
                  <div
                    key={`${x}-${y}`}
                    className="w-12 h-12 grid place-content-center"
                  >
                    <input
                      type="text"
                      value={value === 0 ? "" : value}
                      onChange={(e) => updateValue(x, y, e.target.value)}
                      min={0}
                      max={9}
                      className="w-6 bg-black outline-0"
                    />
                  </div>
                );
              })}
            </div>
          );
        })}
      </div>
      <div className="mt-16">
        <button className="rounded-full" onClick={solveGame}>
          can it be solved?
        </button>
        {isSolvable != null && (
          <div>{isSolvable ? "Solvable" : "Unsolvable"}</div>
        )}
      </div>
    </div>
  );
};

export default SudokuGame;
