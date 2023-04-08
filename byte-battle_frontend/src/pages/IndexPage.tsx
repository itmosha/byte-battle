import React from "react";
import "../../src/index.css"
import Header from "../page-components/Header";
import MainPageContent from "../page-components/MainPageContent";
import Footer from "../page-components/Footer";


interface Props {}

const IndexPage: React.FC<{}> = (props: Props): JSX.Element => {
  return (
    <div className="relative h-full w-full">
      <Header />
      <MainPageContent />
      <Footer />
    </div>
  )
}

export default IndexPage