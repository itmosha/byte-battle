import React from "react";
import "../../src/index.css"


interface Props {}

const IndexPage: React.FC<{}> = (props: Props): JSX.Element => {
  return (
    <>
      <h1 className="text-3xl font-bold underline text-cyan-500">
        IndexPage
      </h1>
    </>
  )
}

export default IndexPage