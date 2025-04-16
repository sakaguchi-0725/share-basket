import { client } from "@/shared/api"
import { ApolloProvider } from "@apollo/client"

type Props = {
  children: React.ReactNode
}

export const AppProvider = ({ children }: Props) => {
  return (
    <ApolloProvider client={client}>
      {children}
    </ApolloProvider>
  )
}