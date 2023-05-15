"use client";

import { FC, useState } from "react";

const sizeOfBoard = 9;

const getNewBoard = (size: number) =>
  Array.from(new Array<string>(size), () =>
    Array.from(new Array<string>(size), () => "")
  );

const getNewBoard2 = (size: number) => {
  let board = getNewBoard(size);
  for (let x = 0; x < size; x++) {
    for (let y = 0; y < size; y++) {
      board[x][y] = `${x}${y}`;
    }
  }
  return board;
};

const Board: FC = () => {
  const [board, setBoard] = useState(getNewBoard2(sizeOfBoard));

  const updateValue = (x: number, y: number, value: string) => {
    const newBoard = [...board];
    newBoard[x][y] = value;
    setBoard(newBoard);
  };

  return (
    <div className="[&>*:nth-child(3n+4)]:border-t-4">
      {board.map((row, x) => {
        return (
          <div
            key={x}
            className="grid grid-cols-9 [&>*:nth-child(3n+4)]:border-l-4"
          >
            {row.map((value, y) => {
              return (
                <div key={`${x}-${y}`} className="w-12 h-12">
                  <input
                    type="text"
                    value={value}
                    onChange={(e) => updateValue(x, y, e.target.value)} // TODO: validate input
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
  );
};

export default Board;
