package student

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {

	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
	} else {
		root.Right = BTreeInsertData(root.Right, data)
	}
	if root.Left != nil {
		root.Left.Parent = root
	}
	if root.Right != nil {
		root.Right.Parent = root
	}

	return root
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {

	if root == nil || root.Data == elem {
		return root
	}

	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}
	return BTreeSearchItem(root.Right, elem)

}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {

	if root == nil {
		return
	}

	BTreeApplyInorder(root.Left, f)
	f(root.Data)
	BTreeApplyInorder(root.Right, f)

}

// go-reloaded BSTree-1
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}
	rplc.Parent = node.Parent

	return root
}

// go-reloaded BSTree-2
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	str := []TreeNode{*root}

	for len(str) != 0 {
		f(str[0].Data)

		if str[0].Left != nil {
			str = append(str, *str[0].Left)
		}

		if str[0].Right != nil {
			str = append(str, *str[0].Right)
		}

		str = str[1:]
	}
}

func BTreeMin(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}

	return BTreeMin(root.Left)
}

// go-reloaded BSTree-3
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// если удаляемый узел меньше корневого, уходим в левую ветку корневого узла
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		var temp *TreeNode
		// если у удаляемого узла, только одна ветка
		if root.Left == nil {
			//сохраняем в temp значение правого узла
			temp = root.Right
			return temp
		} else if root.Right == nil {
			//сохраняем в temp значение левого узла
			temp = root.Left
			return temp
		} else {
			// если у удаляемого узла есть 2 ветки
			// находим узел с мин.значением в правой ветке
			temp = BTreeMin(root.Right)
			// копируем значение temp в удаляемый узел
			root.Data = temp.Data
			// удаляем узел с мин.значением
			root.Right = BTreeDeleteNode(root.Right, temp)
		}
	}
	return root
}
