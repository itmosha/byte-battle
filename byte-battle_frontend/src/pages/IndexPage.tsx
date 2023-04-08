import React from "react";
import "../../src/index.css"
import MainPageContent from "../page-components/MainPageContent";
import Layout from "../page-components/Layout";


interface Props {}

const IndexPage: React.FC<{}> = (props: Props): JSX.Element => {
  return (
    <>
      <Layout>
        <MainPageContent />
      </Layout>
    </>
  )
}

export default IndexPage