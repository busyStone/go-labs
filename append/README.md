append
=======

    for i := 0; i < SliceLen; i++ {
      s = append(s, T{Name: fmt.Sprintf("%d", i)})
    }
    

    var t T
      for i := 0; i < SliceLen; i++ {
        t.Name = fmt.Sprintf("%d", i)
        s = append(s, t)
    }
    
  
    for i := 0; i < SliceLen; i++ {
      t := T{Name: fmt.Sprintf("%d", i)}
      s = append(s, t)
    }
    
第一种 objects 最少
