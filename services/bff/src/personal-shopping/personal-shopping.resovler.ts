import { Args, Mutation, Query, Resolver } from '@nestjs/graphql'
import { ShoppingItem } from './dto/shopping-item.dto'
import { CreateShoppingItemInput } from './dto/create.input'
import { PersonalShoppingService } from './personal-shopping.service'

@Resolver()
export class PersonalShoppingResolver {
  constructor(private readonly shoppingService: PersonalShoppingService) {}

  @Query(() => [ShoppingItem])
  getAll(): Promise<ShoppingItem[]> {
    return this.shoppingService.getAll()
  }

  @Mutation(() => ShoppingItem)
  create(@Args('input') input: CreateShoppingItemInput): Promise<ShoppingItem> {
    return this.shoppingService.create(input)
  }
}
