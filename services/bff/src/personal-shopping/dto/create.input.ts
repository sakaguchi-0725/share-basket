import { Field, InputType } from '@nestjs/graphql'

@InputType()
export class CreateShoppingItemInput {
  @Field()
  name: string

  @Field()
  categoryId: number
}
