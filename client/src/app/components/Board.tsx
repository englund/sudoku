"use client";

import { FC, useState } from "react";

const sizeOfBoard = 9;

const getNewBoard = (size: number) =>
  Array.from(new Array<number>(size), () =>
    Array.from(new Array<number>(size), () => 0)
  );

const Board: FC = () => {
  const [board, setBoard] = useState(getNewBoard(sizeOfBoard));

  const updateValue = (x: number, y: number, value: number) => {
    const newBoard = [...board];
    newBoard[x][y] = value;
    setBoard(newBoard);
  };

  return (
    <>
      {board.map((row, x) => {
        return (
          <div key={x}>
            {row.map((col, y) => (
              <input // TODO: make input type text and validate by hand to make it possible for empty values
                key={`${x}-${y}`}
                type=""
                value={col}
                onChange={(e) => updateValue(x, y, +e.target.value)}
                min={0}
                max={9}
                className="w-6 bg-black outline-0"
              />
            ))}
          </div>
        );
      })}
    </>
  );
};

export default Board;
